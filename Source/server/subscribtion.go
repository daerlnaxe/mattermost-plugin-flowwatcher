package main

import(
	"fmt"
	"bytes"
	"context"
	"encoding/json"

)


const SUBSCRIPTIONS_KEY = "flw_subscriptions"



// Subscription Object
type Subscription struct {
	ChannelID	string
	RUUID		string	// UUID
	URL       	string
	TypeOfFlow	FlowType 
	Content   	string
	Links   	[]string
	Timer		uint	// each time we look if we need to post
	NextPost	uint	// not used for the while, decrement until reach 0 and indicate it's time to refresh
	Login		string  // not used for the while
	Password	string  // not used for the while
	IsActive	bool
	//LastPost	string  // not used for the while, show last refresh
}



// Subscriptions map to key value pairs
type Subscriptions struct {
	Subscriptions map[string]*Subscription
}


/*
* Get the permanent Key-Value Set stored by Mattermost for the plugin
*
* Note init a new subcribtions object if there is nothing stored
*/
func (p *FlowWatcherPlugin) getSubscriptions() (*Subscriptions, error) {
	//
	var subscriptions *Subscriptions

	// Get key-value set
	value, err := p.API.KVGet(SUBSCRIPTIONS_KEY)
		
	if err != nil {
		p.API.LogError(err.Error())
		return nil, nil
	}


	// Create a new 'Subscriptions' map
	if value == nil {
		subscriptions = &Subscriptions{Subscriptions: map[string]*Subscription{}}
	// Decode the one given
	} else {
		json.NewDecoder(bytes.NewReader(value)).Decode(&subscriptions)
	}



	for _, value := range subscriptions.Subscriptions {
		p.API.LogError("-----" + value.ChannelID + " | " + value.URL)
	}
	
	return subscriptions, nil
}


// --- Get Key
func makeKeyByURL(channelID string, url string)string{
	return	fmt.Sprintf("%s/%s", channelID, url)
}

// ---- Get Key


// Return a pointer to a subscribption if found
func getValue(currentSubscriptions *Subscriptions, key string) (*Subscription) {


	// check if key with this combination channelID and url already exists
	_, subExists := currentSubscriptions.Subscriptions[key]

	if subExists {
		// get the value
		return	currentSubscriptions.Subscriptions[key]
/*

		if err := p.storeSubscriptions(currentSubscriptions); err != nil {
			p.API.LogError(err.Error())
			return err
		}*/
	}

	return nil
}


// Subscribe process the /flw subscribe <channel> <url>
func (p *FlowWatcherPlugin) subscribe(ctx context.Context, channelID string, url string) error {

	sub := &Subscription{
		ChannelID: channelID,
		URL:       url,
		Content:   "",
		//Links: make[]string{},
	}

	sub.URL=url

	var err error=nil

	p.API.LogError("--- new detection")
	tFlow, err:=detectFlowFormat(getURLContent(url))

	sub.TypeOfFlow=tFlow

	//
	p.API.LogError(fmt.Sprintf("Type of flow : %s", tFlow))

	if err!=nil {
		p.API.LogError(err.Error())
		return err
	}

	// get the stored 'Subscription" map (or new)
	currentSubscriptions, err := p.getSubscriptions()	
	
	if err != nil {
		p.API.LogError(err.Error())
		return err
	}

	// Create key made on channelID and url
	key := makeKeyByURL(channelID, url)

	// get value
	subValue := getValue(currentSubscriptions, key)

	// If it doesn't exist we can add
	if (subValue == nil){
		// We add informations in the subscribtion
		currentSubscriptions.Subscriptions[key] = &Subscription{ChannelID: sub.ChannelID, URL: sub.URL}

		// Writing the subscribtions map
		err = p.storeSubscriptions(currentSubscriptions)
		if err != nil {
			p.API.LogError(err.Error())
			return err
		}
	}


	return nil
}

func (p *FlowWatcherPlugin) unsubscribe(channelID string, url string) error {
	p.API.LogDebug(fmt.Sprintf("unsubscribe  '%s'  from %s channelID", channelID, url))


	currentSubscriptions, err := p.getSubscriptions()
	p.API.LogDebug(fmt.Sprintf("%s", len(currentSubscriptions.Subscriptions)))

	if err != nil {
		p.API.LogError(err.Error())
		return err
	}

	key := makeKeyByURL(channelID, url)


	p.API.LogDebug("key : "  + key)

	_, ok := currentSubscriptions.Subscriptions[key]
	if ok {
		p.API.LogDebug("key found")
		delete(currentSubscriptions.Subscriptions, key)

		// Update
		if err := p.storeSubscriptions(currentSubscriptions); err != nil {
			p.API.LogError(err.Error())
			return err
		}
	}

	return nil
}


func (p *FlowWatcherPlugin) startFlow(channelID string, url string) error {

	currentSubscriptions, err := p.getSubscriptions()
	if err != nil {
		p.API.LogError(err.Error())
		return err
	}

	key := makeKeyByURL(channelID, url)

	_, value := currentSubscriptions.Subscriptions[key]
	if value {
		currentSubscriptions.Subscriptions[key].IsActive=true

		// Force Flow
		p.forceFlow(channelID, url)

		// Update
		if err := p.storeSubscriptions(currentSubscriptions); err != nil {
			p.API.LogError(err.Error())
			return err
		}
	}

	return nil
}


func (p *FlowWatcherPlugin) stopFlow(channelID string, url string) error {

	currentSubscriptions, err := p.getSubscriptions()
	if err != nil {
		p.API.LogError(err.Error())
		return err
	}

	key := makeKeyByURL(channelID, url)
	_, value := currentSubscriptions.Subscriptions[key]
	if value {
		//
		currentSubscriptions.Subscriptions[key].IsActive=false

		// Update
		if err := p.storeSubscriptions(currentSubscriptions); err != nil {
			p.API.LogError(err.Error())
			return err
		}
	}

	return nil
}


/*
func (p *RSSFeedPlugin) addSubscription(key string, sub *Subscription) error {

	/*
		currentSubscriptions.Subscriptions[key] = &Subscription{ChannelID: sub.ChannelID, URL: sub.URL}
		err = p.storeSubscriptions(currentSubscriptions)
		if err != nil {
			p.API.LogError(err.Error())
			return err
		}

	

	return nil
}
*/

/*
*	
*/
func (p *FlowWatcherPlugin) storeSubscriptions(s *Subscriptions) error {
	b, err := json.Marshal(s)
	if err != nil {
		p.API.LogError(err.Error())
		return err
	}

	p.API.KVSet(SUBSCRIPTIONS_KEY, b)
	return nil
}


/*
*
*/
func (p *FlowWatcherPlugin) updateSubscription(subscription *Subscription) error {
	p.API.LogError("--- UpdateSubscription")
	
	currentSubscriptions, err := p.getSubscriptions()
	if err != nil {
		p.API.LogError(err.Error())
		return err
	}


	key := makeKeyByURL(subscription.ChannelID, subscription.URL)
	
	p.API.LogError("key : " + key)

	_, ok := currentSubscriptions.Subscriptions[key]
	if ok {
		p.API.LogDebug("key found")

		currentSubscriptions.Subscriptions[key] = subscription
		if err := p.storeSubscriptions(currentSubscriptions); err != nil {
			p.API.LogError(err.Error())
			return err
		}
	}
	return nil
}
