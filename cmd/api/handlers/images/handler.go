package images

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/wolke-gallery/api/cmd/api/config"
	"github.com/wolke-gallery/api/cmd/api/utils"
)

func NewImage(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "No `file` key found in form data",
		})
		return
	}

	domain := c.PostForm("domain")

	if domain == "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "No `domain` key found in form data",
		})
		return
	}

	if !utils.CheckIfElementExists(config.Config.Domains, domain) {
		domains := strings.Join(config.Config.Domains, ", ")

		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid domain. Valid are " + domains,
		})
		return
	}

	// TODO: Ideally we would tell the user the max file size in a humanized form
	if file.Size > config.Config.MaxFileSize {
		c.JSON(400, gin.H{
			"success": false,
			"message": "File is too big",
		})
		return
	}

	contentType := file.Header["Content-Type"][0]
	var extension string

	switch contentType {
	case "image/jpeg":
		extension = "jpg"
	case "image/png":
		extension = "png"
	case "image/gif":
		extension = "gif"
	default:
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid file type",
		})
		return
	}

	// NOTE: These are for me, dont mind them
	// 🍖🍗🍘🍙🍚🍛🍜🍝🍞🍟🍠🍡🍢🍣🍤🍥🍦🍧🍨🍩🍪🍫🍬🍭🍮🍯🍰🎊🎋🎌🎍🎎🎏🎑🎒🎓🎠🎡🎢🎣🎤🎥🎦🎧🎨🎩🎪🎬🏀🏁🐀🐁🐂🐃🐄🐅🐆🐇🐈🐉🐊🐋🐌
	// 255 emojis i can use
	// ⌚⌛⏩⏪⏫⏬⏰⏳◽◾☔☕♈♉♊♋♌♍♎♏♐♑♒♓♿⚓⚡⚪⚫⚽⚾⛄⛅⛎⛔⛪⛲⛳⛵⛺⛽✅✊✋✨❌❎❓❔❕❗➕➖➗➰➿⬛⬜⭐⭕🀄🃏🆎🆑🆒🆓🆔🆕🆖🆗🆘🆙🆚🈁🈚🈯🈲🈳🈴🈵🈶🈸🈹🈺🉐🉑🌀🌁🌂🌃🌄🌅🌆🌇🌈🌉🌊🌋🌌🌍🌎🌏🌐🌑🌒🌓🌔🌕🌖🌗🌘🌙🌚🌛🌜🌝🌞🌟🌠🌭🌮🌯🌰🌱🌲🌳🌴🌵🌷🌸🌹🌺🌻🌼🌽🌾🌿🍀🍁🍂🍃🍄🍅🍆🍇🍈🍉🍊🍋🍌🍍🍎🍏🍐🍑🍒🍓🍔🍕🍖🍗🍘🍙🍚🍛🍜🍝🍞🍟🍠🍡🍢🍣🍤🍥🍦🍧🍨🍩🍪🍫🍬🍭🍮🍯🍰🍱🍲🍳🍴🍵🍶🍷🍸🍹🍺🍻🍼🍾🍿🎀🎁🎂🎃🎄🎅🎆🎇🎈🎉🎊🎋🎌🎍🎎🎏🎐🎑🎒🎓🎠🎡🎢🎣🎤🎥🎦🎧🎨🎩🎪🎫🎬🎭🎮🎯🎰🎱🎲🎳🎴🎵🎶🎷🎸🎹🎺🎻🎼🎽🎾🎿🏀🏁🏂🏃🏄🏅🏆🏇🏈🏉🏊🏏🏐🏑🏒🏓🏠🏡🏢🏣🏤🏥🏦🏧🏨🏩🏪🏫🏬🏭🏮🏯🏰🏴🏸🏹🏺🏻🏼🏽🏾🏿🐀🐁🐂🐃🐄🐅🐆🐇🐈🐉🐊🐋🐌🐍🐎🐏🐐🐑🐒🐓🐔🐕🐖🐗🐘🐙🐚🐛🐜🐝🐞🐟🐠🐡🐢🐣🐤🐥🐦🐧🐨🐩🐪🐫🐬🐭🐮🐯🐰🐱🐲🐳🐴🐵🐶🐷🐸🐹🐺🐻🐼🐽🐾👀👂👃👄👅👆👇👈👉👊👋👌👍👎👏👐👑👒👓👔👕👖👗👘👙👚👛👜👝👞👟👠👡👢👣👤👥👦👧👨👩👪👫👬👭👮👯👰👱👲👳👴👵👶👷👸👹👺👻👼👽👾👿💀💁💂💃💄💅💆💇💈💉💊💋💌💍💎💏💐💑💒💓💔💕💖💗💘💙💚💛💜💝💞💟💠💡💢💣💤💥💦💧💨💩💪💫💬💭💮💯💰💱💲💳💴💵💶💷💸💹💺💻💼💽💾💿📀📁📂📃📄📅📆📇📈📉📊📋📌📍📎📏📐📑📒📓📔📕📖📗📘📙📚📛📜📝📞📟📠📡📢📣📤📥📦📧📨📩📪📫📬📭📮📯📰📱📲📳📴📵📶📷📸📹📺📻📼📿🔀🔁🔂🔃🔄🔅🔆🔇🔈🔉🔊🔋🔌🔍🔎🔏🔐🔑🔒🔓🔔🔕🔖🔗🔘🔙🔚🔛🔜🔝🔞🔟🔠🔡🔢🔣🔤🔥🔦🔧🔨🔩🔪🔫🔬🔭🔮🔯🔰🔱🔲🔳🔴🔵🔶🔷🔸🔹🔺🔻🔼🔽🕋🕌🕍🕎🕐🕑🕒🕓🕔🕕🕖🕗🕘🕙🕚🕛🕜🕝🕞🕟🕠🕡🕢🕣🕤🕥🕦🕧🕺🖕🖖🖤🗻🗼🗽🗾🗿😀😁😂😃😄😅😆😇😈😉😊😋😌😍😎😏😐😑😒😓😔😕😖😗😘😙😚😛😜😝😞😟😠😡😢😣😤😥😦😧😨😩😪😫😬😭😮😯😰😱😲😳😴😵😶😷😸😹😺😻😼😽😾😿🙀🙁🙂🙃🙄🙅🙆🙇🙈🙉🙊🙋🙌🙍🙎🙏🚀🚁🚂🚃🚄🚅🚆🚇🚈🚉🚊🚋🚌🚍🚎🚏🚐🚑🚒🚓🚔🚕🚖🚗🚘🚙🚚🚛🚜🚝🚞🚟🚠🚡🚢🚣🚤🚥🚦🚧🚨🚩🚪🚫🚬🚭🚮🚯🚰🚱🚲🚳🚴🚵🚶🚷🚸🚹🚺🚻🚼🚽🚾🚿🛀🛁🛂🛃🛄🛅🛌🛐🛑🛒🛕🛖🛗🛫🛬🛴🛵🛶🛷🛸🛹🛺🛻🛼🟠🟡🟢🟣🟤🟥🟦🟧🟨🟩🟪🟫🤌🤍🤎🤏🤐🤑🤒🤓🤔🤕🤖🤗🤘🤙🤚🤛🤜🤝🤞🤟🤠🤡🤢🤣🤤🤥🤦🤧🤨🤩🤪🤫🤬🤭🤮🤯🤰🤱🤲🤳🤴🤵🤶🤷🤸🤹🤺🤼🤽🤾🤿🥀🥁🥂🥃🥄🥅🥇🥈🥉🥊🥋🥌🥍🥎🥏🥐🥑🥒🥓🥔🥕🥖🥗🥘🥙🥚🥛🥜🥝🥞🥟🥠🥡🥢🥣🥤🥥🥦🥧🥨🥩🥪🥫🥬🥭🥮🥯🥰🥱🥲🥳🥴🥵🥶🥷🥸🥺🥻🥼🥽🥾🥿🦀🦁🦂🦃🦄🦅🦆🦇🦈🦉🦊🦋🦌🦍🦎🦏🦐🦑🦒🦓🦔🦕🦖🦗🦘🦙🦚🦛🦜🦝🦞🦟🦠🦡🦢🦣🦤🦥🦦🦧🦨🦩🦪🦫🦬🦭🦮🦯🦰🦱🦲🦳🦴🦵🦶🦷🦸🦹🦺🦻🦼🦽🦾🦿🧀🧁🧂🧃🧄🧅🧆🧇🧈🧉🧊🧋🧍🧎🧏🧐🧑🧒🧓🧔🧕🧖🧗🧘🧙🧚🧛🧜🧝🧞🧟🧠🧡🧢🧣🧤🧥🧦🧧🧨🧩🧪🧫🧬🧭🧮🧯🧰🧱🧲🧳🧴🧵🧶🧷🧸🧹🧺🧻🧼🧽🧾🧿🩰🩱🩲🩳🩴🩸🩹🩺🪀🪁🪂🪃🪄🪅🪆🪐🪑🪒🪓🪔🪕🪖🪗🪘🪙🪚🪛🪜🪝🪞🪟🪠🪡🪢🪣🪤🪥🪦🪧🪨🪰🪱🪲🪳🪴🪵🪶🫀🫁🫂🫐🫑🫒🫓🫔🫕🫖©️®️‼️⁉️™️ℹ️↔️↕️↖️↗️↘️↙️↩️↪️⌨️⏏️⏭️⏮️⏯️⏱️⏲️⏸️⏹️⏺️Ⓜ️▪️▫️▶️◀️◻️◼️☀️☁️☂️☃️☄️☎️☑️☘️☝️☠️☢️☣️☦️☪️☮️☯️☸️☹️☺️♀️♂️♟️♠️♣️♥️♦️♨️♻️♾️⚒️⚔️⚕️⚖️⚗️⚙️⚛️⚜️⚠️⚧️⚰️⚱️⛈️⛏️⛑️⛓️⛩️⛰️⛱️⛴️⛷️⛸️⛹️✂️✈️✉️✌️✍️✏️✒️✔️✖️✝️✡️✳️✴️❄️❇️❣️❤️➡️⤴️⤵️⬅️⬆️⬇️〰️〽️㊗️㊙️🅰️🅱️🅾️🅿️🈂️🈷️🌡️🌤️🌥️🌦️🌧️🌨️🌩️🌪️🌫️🌬️🌶️🍽️🎖️🎗️🎙️🎚️🎛️🎞️🎟️🏋️🏌️🏍️🏎️🏔️🏕️🏖️🏗️🏘️🏙️🏚️🏛️🏜️🏝️🏞️🏟️🏳️🏵️🏷️🐿️👁️📽️🕉️🕊️🕯️🕰️🕳️🕴️🕵️🕶️🕷️🕸️🕹️🖇️🖊️🖋️🖌️🖍️🖐️🖥️🖨️🖱️🖲️🖼️🗂️🗃️🗄️🗑️🗒️🗓️🗜️🗝️🗞️🗡️🗣️🗨️🗯️🗳️🗺️🛋️🛍️🛎️🛏️🛠️🛡️🛢️🛣️🛤️🛥️🛩️🛰️🛳️#️⃣*️⃣0️⃣1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣
	id, err := gonanoid.New(config.Config.IdLength)

	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Failed to generate id.. please try again",
		})
		return
	}

	name := fmt.Sprintf("%s.%s", id, extension)

	if err := c.SaveUploadedFile(file, config.Config.Directory+name); err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Failed to save file",
		})
		return
	}

	url := fmt.Sprintf("https://%s/%s", domain, name)

	c.JSON(200, gin.H{
		"success": true,
		"message": url,
	})
}
