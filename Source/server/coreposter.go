package main

import(
	"fmt"
	
	"time"
	"github.com/mattermost/mattermost/server/public/model"
	// RSS/ATOM Parser
	"github.com/mmcdole/gofeed"
	
	// Convert HTML To Markdown
	"github.com/JohannesKaufmann/html-to-markdown"
	"github.com/JohannesKaufmann/html-to-markdown/plugin"
)


const (
	DEBUG = true
)

// First to Launch Core
func (p *FlowWatcherPlugin) initCorePoster() {
	wakeUpTime, err := p.getWakeUpTime()
	
	if err != nil {
		p.API.LogError(err.Error())
	}


	// golang seems to not implement a while loop, for works.
	for p.corePosterFlag {
		
		// Run poster manager
		err := p.subscribtionManager()
		

		if err != nil {
			p.API.LogError(err.Error())

		}

		// Impossible to fall under a minute (for the)
		// Mandatory to use time.Duration() don't accept int * minutes
		time.Sleep(time.Duration( wakeUpTime) * time.Minute)
	}
}



func (p *FlowWatcherPlugin) subscribtionManager()error{
	// get the stored 'Subscription" map (or new)
	currentSubscriptions, err := p.getSubscriptions()
	if err != nil {
		p.API.LogError(err.Error())
		return err
	}


	
	for _, value := range currentSubscriptions.Subscriptions {
		//err := p.createBotPost(value.ChannelID, "I'm still standing yeah yeah yeah", "")
		//err := p.parseContent(value)

		// timeout
		/*ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()*/
		// timeout


		//feed, err := fp.ParseURL("http://feeds.twit.tv/twit.xml")
		//feed, err := fp.ParseURL("https://www.jeuxvideo.com/rss/rss.xml")//, ctx)

		if !value.IsActive{
			return nil
		}

		p.getFlow(value)

	}
	return nil
}


// ForceFlow above timer
func (p *FlowWatcherPlugin) forceFlow(channelID string, url string) error {
	currentSubscriptions, err := p.getSubscriptions()
	if err != nil {
		p.API.LogError(err.Error())
		return err
	}

	key := makeKeyByURL(channelID, url)
	_, ok := currentSubscriptions.Subscriptions[key]
	if ok {
		p.getFlow(currentSubscriptions.Subscriptions[key])
	}

	return nil

}


// Common method to get flow, update and send to channel
func (p *FlowWatcherPlugin) getFlow(value *Subscription){
	
	p.API.LogError("Parsing: '"+value.URL +"'")
	//fp := &gofeed.NewParser{}

	//if (value.TypeOfFlow == XML){
		fp := gofeed.NewParser()		
		
	//}else if (value.TypeOfFlow == JSON){
//		fp := json.Parser{}
	//	fp = gofeed.NewParser()		

	//}
	feed, err := fp.ParseURL(value.URL )
		

	//
	if err != nil {
		value.IsActive=false
		p.API.LogError(">>>> Parsing Error"+ err.Error())
		p.updateSubscription(value)		

		return
	}


	//
	if DEBUG {
		fmt.Println(feed.Title)			
	}

	// Storing hash of urls
	old_links := value.Links[:]
	new_links := []string{}
	newPostFound :=0

	p.API.LogError(fmt.Sprintf(">>> value: %d | old: %d | new: %d | newtmp: %d", len(value.Links), len(old_links), len(feed.Items), len(new_links)))

	// Adding only existing
	for _, item := range feed.Items{
		p.API.LogError(fmt.Sprintf(">>> Research new %s", item.Link))

		found := false
		// Removing Old by not adding them
		for _, old_link := range old_links{
			p.API.LogDebug(">>> Research old " + old_link)
			if(old_link == item.Link){
				p.API.LogError(">>> Found old " + old_link)
				found = true	
				break
			// Non présent <- on verra pour optimiser
			}
		}
		
		// 100% new -> Post
		if !found{
			newPostFound++
			// Fonctionne sans image si j'ai bien compris (enfin si, c'est compliqué)
			p.sendItem(value.ChannelID, item);		
		}
		// Adding in all cases
		new_links = append(new_links, item.Link)
	}

	
	// Store new result in subscribption
	value.Links = new_links[:]
	p.API.LogError(fmt.Sprintf(">>> value: %d | old: %d | new: %d | newtmp: %d", len(value.Links), len(old_links), len(feed.Items), len(new_links)))
	
	p.API.LogError(fmt.Sprintf("Number of fresh news: %d", newPostFound))

	if(newPostFound>0 ){
	
		p.updateSubscription(value)		
	}
}






// Sent Item to channel
func (p *FlowWatcherPlugin) sendItem(channelID string, item *gofeed.Item) {
	
	// !!!! Rajouter un comparateur pour ne pas reposter les mêmes messages
	converter := md.NewConverter("", true, nil)

		// Use the `GitHubFlavored` plugin from the `plugin` package.
		converter.Use(plugin.GitHubFlavored())

	
		if DEBUG{
			/*
			fmt.Println("--------------------------------------------------------------------------------------------------------------")

			fmt.Println(value.Title)
			//fmt.Println(value.Description)
			//fmt.Println(value.Content)
			/*fmt.Println("Enclosures")
			fmt.Println(value.Enclosures)
			fmt.Println("Links")
			fmt.Println(value.Links)
			//fmt.Println("Image")
			//fmt.Println(value.Image)
			fmt.Println("--------------------------------------------------------------------------------------------------------------")		
			*/
		}

		  

		/*
		if config.FormatTitle {
			message ="##### "
		}*/
		
		//message += fmt.Sprintf("[%s](%s)\n",value.Link,value.Link)

		fmt.Println(fmt.Sprintf("Title (taille): %d", len(item.Title)))
		
		fmt.Println(fmt.Sprintf("Description (taille): %d", len(item.Description)))
		fmt.Println(fmt.Sprintf("Content (taille): %d", len(item.Content)))

		fmt.Println(fmt.Sprintf("URLImage: %s", item.Image.URL))



		// Title
		message := "##### "
		//message += item.Title + "\n"
		// Title


		// Link
		message += fmt.Sprintf("[%s](%s) \n",item.Title,item.Link)

		// Summary
		message += "\n**Summary**:\n"
		/*

			if config.FormatTitle {
				post = post + "###### "
			}
			post = post + item.Title + "\n"
		}*/
	
			/*
		if config.ShowRSSLink {
			post = post + strings.TrimSpace(item.Link) + "\n"
		}
			*/
		
		/*
			if config.ShowDescription {
			post = post + html2md.Convert(item.Description) + "\n"
		}
		*/
		
		// Content
		p.API.LogDebug("Try Convert string to md")

		markdown, err := converter.ConvertString(item.Description)
			

		if err != nil {
			p.API.LogError("%s",err)
		  }

		
		if(len(markdown) > 200){
			p.API.LogDebug(fmt.Sprintf("cut markdown: %s", len(markdown) ))
			message += markdown[:200]+ "[...]\n\n"
		}
		// Content

		// image
		if(len(item.Image.URL)>0){
			message+=fmt.Sprintf("![%s](%s =200)", item.Image.Title,  item.Image.URL)
		}
		//


		// Published
		p.API.LogDebug("Published: ")
		message+= fmt.Sprintf("*%s", item.Published)
		// Published

		// Author
		p.API.LogDebug("Author:")
		if(item.Author != nil){
			message +=  fmt.Sprintf(" by %s*\n", item.Author.Name)
		}
		// Autor
		

		//message += value.Link

		//message += fmt.Sprintf( "![%s](%s \"Mattermost Icon\")",value.Image.Title, value.Image.URL)
		//message += fmt.Sprintf( "![%s](%s =200 \"Mattermost Icon\")",value.Image.Title, value.Image.URL)
//		message += "----mwahahhaha\n----"
/*
		message += fmt.Sprintf( "![%s](%s =200 \"Hover text\")",value.Image.Title, value.Image.URL)
		message += "----mwahahhaha\n----"
*/
		
///		message += "-----\r\n"
		/*markdown2, err := converter.ConvertString(value.Content)
		message+=markdown2*/

		p.createBotPost(channelID, message, "", item.Image.URL)
	}






func (p *FlowWatcherPlugin) createBotPost(channelID string, message string, postType string, image string) error {
/*
	mee :=`json:"attachments": [
        {
            "fallback": "test",
            "color": "#FF8000",
            "pretext": "This is optional pretext that shows above the attachment.",
            "text": "This is the text of the attachment. It should appear just above an image of the Mattermost logo. The left border of the attachment should be colored orange, and below the image it should include additional fields that are formatted in columns. At the top of the attachment, there should be an author name followed by a bolded title. Both the author name and the title should be hyperlinks.",
            "author_name": "Mattermost",
            "author_icon": "https://mattermost.com/wp-content/uploads/2022/02/icon_WS.png",
            "author_link": "https://mattermost.org/",
            "title": "Example Attachment",
            "title_link": "https://developers.mattermost.com/integrate/reference/message-attachments/",
            "fields": [
                {
                    "short":false,
                    "title":"Long Field",
                    "value":"Testing with a very long piece of text that will take up the whole width of the table. And then some more text to make it extra long."
                },
                {
                    "short":true,
                    "title":"Column One",
                    "value":"Testing"
                },
                {
                    "short":true,
                    "title":"Column Two",
                    "value":"Testing"
                },
                {
                    "short":false,
                    "title":"Another Field",
                    "value":"Testing"
                }
            ],
            "image_url": "https://mattermost.com/wp-content/uploads/2022/02/icon_WS.png"
        }
    ]
	}
	`
*/
		
	post := &model.Post{
		UserId:    p.botUserID,
		ChannelId: channelID,
		Message:   message,
		//Type:      "custom_git_pr",
		/*Props: map[string]interface{}{
			"from_webhook":      "true",
			"override_username": botDisplayName,
		},*/
		
	}
	
	/*draft := &model.Draft{ 
		Message: image,
	}*/
	
	draft := map[string]any{image: image}  /* &model.StringInterface [
		{image}: image
	]*/
	//draft[image]= image

	post.SetProps(draft)


	if _, err := p.API.CreatePost(post); err != nil {
		p.API.LogError(
			"We could not create the RSS post",
			"user_id", post.UserId,
			"err", err.Error(),
		)
	}

	return nil
}
