## Dialog Message
### model.DialogElement

|Field|Type|Description|
|-|-|-|
|DisplayName|string|Display name of the field shown to the user in the dialog. Maximum 24 characters.|
|Name|string|Name of the field element used by the integration. Maximum 300 characters. You should use UNIQUE name fields in the same dialog|
|Type|string|text, texareat, select, bool, radio|
|SubType|string|text & textarea:	text, email, number, password (as of v5.14), tel, or url.|
|Default|string|(Optional) Set a default value for this form element. |
|Placeholder|string|(Optional) A string displayed to include a label besides the element. (max 150)|
|HelpText|string|(Optional) Set help text for this form element.(max 150)|
|Optional|Boolean|(Optional) Set to true if this form element is not required. Default is false.|

MinLength   				// int      			(Optional) Minimum input length allowed for an element. Default is 0.
MaxLength					// int					(Optional) Maximum input length allowed for an element. Default is 3000.
DataSource  				// string 				(Optional) One of users, or channels. If none specified, assumes a manual list of options is provided by the integration.
Options     				//[]*PostActionOptions	(Optional) An array of options for the select element. Not applicable for users or channels data sources.

<br>
<br>


### Résultat du Submit
```json
{
    "callback_id": "",
    "cancelled": false,
    "channel_id": "ej11b7f5qbrodyk7u3x8yhqzko",
    "state": "",
    "submission": {
//--- custom
        "login": "r",
        "pass": "raz",
        "setactive": "act",
        "urlLink": "https://www.jeuxvideo.com/rss/rss.xml"
//--- custom
    },
    "team_id": "jr74qu6sbidgmkyitqsiwcohph",
    "type": "dialog_submission",
    "user_id": "ujtnkspj788bdqi45ehfpufhfo"
}
```

<br>
<br>

### Traitement
#### Interface générique
```golang
// Conversion du body qui a été convertis du json
m := rst.(map[string]interface{})

// Affichage du champ user_id
p.API.LogDebug(fmt.Sprintf("test: %s",m["user_id"]))	


// Récupération de la partie submission qui est une struct
form := m["submission"]
// Affichage
p.API.LogDebug(fmt.Sprintf("test form: %s", form))
// Conversion de la partie submission vers un map string/interface	
formCont:=form.(map[string]interface{})


p.API.LogDebug(fmt.Sprintf("test urlLink: %s", formCont["urlLink"]))
p.API.LogDebug(fmt.Sprintf("test setactive: %s", formCont["setactive"]))
p.API.LogDebug(fmt.Sprintf("test login: %s", formCont["login"]))
p.API.LogDebug(fmt.Sprintf("test pass: %s", formCont["pass"]))
```

<br>
<br>

### Tips
#### Json
Si la structure a unchamp qui ne comporte pas une majuscule, ça ne fonctionne pas. Il faut mettre à droite du type `json: <nom json>`
