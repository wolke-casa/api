package images

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/wolke-gallery/api/config"
	"github.com/wolke-gallery/api/database/models"
	"github.com/wolke-gallery/api/medium"
	"github.com/wolke-gallery/api/utils"
)

/*
curl -X POST http://localhost:8080/images/new \
-H 'Content-Type: multipart/form-data' \
-H 'Authorization: owo' \
-F 'domain=domiscute.com' \
-F 'file=@/home/dominic/Downloads/33_left.jpg'
*/

func GetImage(c *gin.Context) {
	var data models.RequestGetImage

	if err := c.ShouldBindUri(&data); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "No `id` in uri",
		})
		return

	}

	reader, err := medium.Storage.Get(data.Id)

	if err != nil {
		fmt.Println(err)

		c.JSON(404, gin.H{
			"success": false,
			"message": "That image doesnt exist",
		})
		return
	}

	bytes := utils.IoReaderToByteSlice(reader)

	bytes512 := make([]byte, 512)
	copy(bytes512, bytes)

	contentType := http.DetectContentType(bytes512)

	if err != nil {
		fmt.Println(err)

		c.JSON(500, gin.H{
			"success": false,
			"message": "Unknown error occurred",
		})
		return
	}

	fmt.Println(contentType)

	c.Data(200, contentType, bytes)
}

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

		c.JSON(404, gin.H{
			"success": false,
			"message": "Invalid domain. Valid are " + domains,
		})
		return
	}

	// TODO: Ideally we would tell the user the max file size in a humanized form
	if file.Size > config.Config.MaxFileSize {
		c.JSON(413, gin.H{
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

	src, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Failed to get file",
		})
		return
	}
	defer src.Close()

	err = medium.Storage.Put(src, name)

	if err != nil {
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
