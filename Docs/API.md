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


Optional    				// Boolean				(Optional) Set to true if this form element is not required. Default is false.		
MinLength   				// int      			(Optional) Minimum input length allowed for an element. Default is 0.
MaxLength					// int					(Optional) Maximum input length allowed for an element. Default is 3000.
DataSource  				// string 				(Optional) One of users, or channels. If none specified, assumes a manual list of options is provided by the integration.
Options     				//[]*PostActionOptions	(Optional) An array of options for the select element. Not applicable for users or channels data sources.
