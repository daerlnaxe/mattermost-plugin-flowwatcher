package main

/*
Commands list used

Called by activate.go

*/

import (
	"context"
	"fmt"
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
	//"github.com/mmcdole/gofeed"

	
//	"github.com/mattermost/mattermost/server/public/shared/mlog"
	"strings"
)


// The help you see when you type /flw {*}
const COMMAND_HELP = `* |/flw sub {url}| - Connect your Mattermost channel to a flow.
* |/flw ls| - List RSS Subscribtions for this Mattermost channel.
* |/flw rem| - Remove a flow from your Mattermost channel.
* |/flw act {flow}| Activating a flow for this Mattermost channel.
* |/flw stop {flow}| Disabling a flow for this Mattermost channel.
* |/flw lsChan [UserId]| - List ALL posts by FlowWatcher bot (Optionnal someone else)
* |/flw clrchan [UserId]| - Remove ALL posts by FlowWatcher bot (Optionnal someone else).
`


// No command can work without it < ? sure ?
const (
	commandTriggerSub       = "sub"
	commandTriggerRem       = "rem"
	commandTriggerAct       = "act"
	commandTriggerStop      = "stop"
	commandTriggerList      = "ls"
	commandTriggerHelp      = "help"
	CommandTriggerlsChan	= "lschan"
	CommandTriggerClrChan	= "clrchan"
)


// ---- API

// Commands description
// API
func getCommand() *model.Command {
	return &model.Command{
		// Slack command "/flw"
		Trigger:          "flw",
		DisplayName:      "flowwatcher",
		Description:      "Allows user to subscribe to a flow.",
		AutoComplete:     true,
		AutoCompleteDesc: "Available commands:  help, ls, sub, rem, act, stop,lschan,clrchan",
		AutoCompleteHint: "[command]",
		AutocompleteData: getAutocompleteData(),
	}
}


/*
	Caution: cname of Autocomplete determinate the name of the slack command
	You can't see commands in command if not referenced correctly here

	- var
	- string
	- addcommand
*/
func getAutocompleteData() *model.AutocompleteData {
	flw := model.NewAutocompleteData("flw", "[command]",
		"Available commands: help, sub, ls, clrchan, rmpost")
	// subscribe		
	sub := model.NewAutocompleteData("sub", "", "Add a Flow to watch")
	flw.AddCommand(sub)
	// list flows
	ls := model.NewAutocompleteData("ls", "", "List Flow for the channel")
	flw.AddCommand(ls)	
	// unsubscribe		
	rem := model.NewAutocompleteData("rem", "", "Add a Flow to watch")
	flw.AddCommand(rem)
	// activate flow
	act := model.NewAutocompleteData("act", "", "Add a Flow to watch")
	flw.AddCommand(act)
	// disabling flow
	stop := model.NewAutocompleteData("stop", "", "Add a Flow to watch")
	flw.AddCommand(stop)
	// list posts on chan
	lschan := model.NewAutocompleteData("lschan", "", "List ALL messages from FlowWatcher Bot [optionanl: userid]")
	flw.AddCommand(lschan)
    // clear posts on chan
	clrchan := model.NewAutocompleteData("clrchan", "", "Clear ALL messages from FlowWatcher Bot [optionanl: userid]")
	flw.AddCommand(clrchan)


		
	return flw
}


// Check  url
func verifyURL(parameters []string)(bool, string){
	if len(parameters) == 0 {
		return false,"Please specify a url."
	} else if len(parameters) > 1 {
		return false,"Please specify a valid url."
	}

	return true,""
}



// Will interact with user when he is using Slack commands
// Its here to specify all commands
// API
func (p *FlowWatcherPlugin) ExecuteCommand(_ *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	// Split for command args
	split := strings.Fields(args.Command)
	
	command := split[0]
	parameters := []string{}
	action := ""

	// No param
	if len(split) > 1 {
		action = split[1]
	}

	// With params
	if len(split) > 2 {
		parameters = split[2:]
	}

	// Doesn't concern flwatcher => pass
	if command != "/flw" {
		return &model.CommandResponse{}, nil
	}



	// ---- Administrators only
	switch action{
		case "lschan","clrchan":
		isSysadmin, _ := p.hasSysadminRole(args.UserId)

		if (isSysadmin == false){
			//return getCommandResponse("ephemeral", "You are not admin my dear ..."), nil
			return getCommandResponse("ephemeral", "You are not admin my dear, mon chéri, meine Liebe, hon hon hon ..."), nil
		}
	}


	// Managing actions
	switch action {
	// help
	case "help":
		text := "###### Mattermost FlowWatcher Plugin - Slash Command Help\n" + strings.Replace(COMMAND_HELP, "|", "`", -1)
		return getCommandResponse("ephemeral", text), nil

	// Subscribing to a Rss feed
	case "sub":
		res, msg := verifyURL( parameters)
		if !res{
			return getCommandResponse("ephemeral", msg), nil
		}

		url := parameters[0]


		/*
		// Test to see if it's valid
		fp := gofeed.NewParser()		
		_, err := fp.ParseURL(url )
	
		//
		if err != nil {
			return getCommandResponse("ephemeral", err.Error()), nil
		}*/

		// Try to register to the Flow
		if err := p.subscribe(context.Background(), args.ChannelId, url); err != nil {
			return getCommandResponse("ephemeral", err.Error()), nil
		}
		
		// Return this message if succesfull
		return getCommandResponse("ephemeral", fmt.Sprintf("Successfully subscribed to %s.", url)), nil

	// list of subscribtions for the channel
	case "ls"	:
		
		subscriptions, err := p.getSubscriptions()

		if err != nil {
			return getCommandResponse("ephemeral", err.Error()), nil
		}

		// Check Subscribtions
		index := 0
		txt := "### Flow Subscriptions in this channel\n"

		//commits := map[Subscription]int{}
		

		/*for i := 0; i < 10; i++ {
			value := subscriptions.Subscriptions[i]

			if value.ChannelID == args.ChannelId {
				index ++
				txt += fmt.Sprintf("* %s) `%s`\n", index,value.URL)
			}			
			
			
		}*/
		
		for _, value := range subscriptions.Subscriptions {
			
			if value.ChannelID == args.ChannelId {
				txt += fmt.Sprintf("* `%d`) %s | ", index+1,value.URL)
				if value.IsActive{
					txt+= "Active"
				}else{
					txt+= "Disabled" 
				}
				
				txt +="`\n"
				index++
			}			
			
		}

		// If no subscription, manage differently
		if(index==0){
			return getCommandResponse("ephemeral", "There is no subscribtion on this channel"), nil
		// There is entries
		}else{
			return getCommandResponse("ephemeral", txt), nil
		}

	case "rem":
		p.API.LogError("rem")
		res, msg := verifyURL( parameters)
		if !res{ 
			return getCommandResponse("ephemeral", msg), nil
		}

		url := parameters[0]

		// Try to unregister the Flow
		if err := p.unsubscribe(args.ChannelId, url); err != nil {
			return getCommandResponse("ephemeral", err.Error()), nil
		}
		
		// Return this message if succesfull
		return getCommandResponse("ephemeral", fmt.Sprintf("Successfully unsubscribed from %s.", url)), nil
	// Activate a flow
    case "act":
		res, msg := verifyURL( parameters)
		if !res{
			return getCommandResponse("ephemeral", msg), nil
		}

		url := parameters[0]

		// Try to Activate the Flow
		if err := p.startFlow(args.ChannelId, url); err != nil {
			return getCommandResponse("ephemeral", err.Error()), nil
		}
		
		// Return this message if succesfull
		return getCommandResponse("ephemeral", fmt.Sprintf("Successfully activate flow '%s'.", url)), nil
	// Stop a flow
	case "stop":
		res, msg := verifyURL( parameters)
		if !res{
			return getCommandResponse("ephemeral", msg), nil
		}

		url := parameters[0]

		// Try to Stop the Flow
		if err := p.stopFlow(args.ChannelId, url); err != nil {
			return getCommandResponse("ephemeral", err.Error()), nil
		}
		
		// Return this message if succesfull
		return getCommandResponse("ephemeral", fmt.Sprintf("Successfully stopped flow '%s'.", url)), nil

	// Will move to an another bot
	case "lschan":		

		userID := ""
		who := &model.User{}
		
		// Case no UserId specified we will delete bot message
		if(len(parameters)==0){			
			userID=p.botUserID			
			who,_ = p.API.GetUser(userID)
		// More than 1 user id
		}else if (len(parameters)>1){
			return getCommandResponse("ephemeral", "Sorry but there is too much parameter"), nil	
		}else{
			userID= parameters[0]
			who,_ = p.API.GetUserByUsername(userID)
			
		}
		p.listInChanel(args.ChannelId, who)//, value.Team)
	
		return getCommandResponse("ephemeral", fmt.Sprintf("List messages for %s performed.", who.Username)), nil

	// Will move to another bot
    case "clrchan":
		p.API.LogDebug("---- > Test ClrChan")
		//subscriptions, err := p.getSubscriptions()

		userID := ""
		who := &model.User{}


		// Case no UserId specified we will delete bot message
		if(len(parameters)==0){			
			userID=p.botUserID			
			who,_ = p.API.GetUser(userID)			
		// More than 1 user id
		}else if (len(parameters)>1){
			return getCommandResponse("ephemeral", "Sorry but there is too much parameter"), nil	
		}else{
			userID= parameters[0]
			who,_ = p.API.GetUserByUsername(userID)
			
		}

		//who,_ := p.API.GetUser(userID)
		
		p.cleanChannel(args.ChannelId, userID)//, value.Team)
	
		return getCommandResponse("ephemeral", fmt.Sprintf("Cleaning up %s's performed.", who.Username)), nil

	// Default
	default:
		text := "###### Mattermost FlowWatcher Plugin - Slash Command Help\n" + strings.Replace(COMMAND_HELP, "|", "`", -1)
		return getCommandResponse("ephemeral", text), nil
	}

	return nil ,nil

}


// ---- END API



/* Le type CommandResponse a besoin maintenant que type soit une string.Add a RSS to watch
	- canal: in_channel
*/
func getCommandResponse(responseType, text string) *model.CommandResponse {
	return &model.CommandResponse{
		ResponseType: responseType,
		Text:         text,
		Username:     BOT_DISPLAY_NAME,
		//IconURL:     ,
		Type:         "ephemeral",
	}
}



// Check if user has admin role
func (p *FlowWatcherPlugin) hasSysadminRole(userID string) (bool, error) {
	user, appErr := p.API.GetUser(userID)
	if appErr != nil {
		return false, appErr
	}
	if !strings.Contains(user.Roles, "system_admin") {
		return false, nil
	}
	return true, nil
}


// ---- Process
func (p *FlowWatcherPlugin)listInChanel(channelId string, user *model.User) error {
	mee,_ := p.API.GetPostsForChannel(channelId, 0, 2000 )


	p.API.LogError(fmt.Sprintf("listChannel: %s", channelId))


	for _, value := range mee.Posts{

		if(value.UserId == user.Id){
			p.API.LogError( fmt.Sprintf(">>> post: %s | par: %s (%s) | %s", value.Id, user.Username, value.UserId, value.Message))			
		}
	}

	// Refresh posts for the channel

	p.API.GetPostsForChannel(channelId, 0, 2000) 


	return nil
}



func (p *FlowWatcherPlugin)cleanChannel(channelId string, UserId string) error {
	mee,_ := p.API.GetPostsForChannel(channelId, 0, 500 )


	p.API.LogError(fmt.Sprintf("clearchannel: %s", channelId))

	
	for _, value := range mee.Posts{
		deletePost := false
		
		if(value.UserId == UserId){
			deletePost = true	
		}

		// Pas de suppression si taggé comme validé
		if(value.HasReactions && value.UserId == p.botUserID ){
			reactions,_ := p.API.GetReactions(value.Id) 
			for _, reaction := range reactions{
				switch (reaction.EmojiName){
				case "white_check_mark" :
					deletePost=false
				}
			}


		}

		if deletePost{
			p.API.LogError( fmt.Sprintf(">>> delete post: %s | par: %s", value.Id, value.UserId))
			p.API.DeletePost(value.Id)
		}
	}


	

	return nil
}

// ---- Process