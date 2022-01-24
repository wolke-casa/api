package images

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/sbani/go-humanizer/units"
	"github.com/wolke-gallery/api/config"
	"github.com/wolke-gallery/api/database/models"
	"github.com/wolke-gallery/api/handlers"
	"github.com/wolke-gallery/api/storage"
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
		error := handlers.ErrMissingData

		c.JSON(error.Status, gin.H{
			"success": false,
			"message": strings.Replace(error.Error, "{}", "id", 1),
		})
		return
	}

	reader, err := storage.Storage.Get(data.Id)

	if err != nil {
		error := handlers.ErrResourceNotFound

		c.JSON(error.Status, gin.H{
			"success": false,
			"message": error.Error,
		})
		return
	}

	bytes := utils.IoReaderToByteSlice(reader)

	bytes512 := make([]byte, 512)
	copy(bytes512, bytes)

	contentType := http.DetectContentType(bytes512)

	if err != nil {
		error := handlers.ErrUnknownErrorOccurred

		c.JSON(error.Status, gin.H{
			"success": false,
			"message": error.Error,
		})
		return
	}

	fmt.Println(contentType)

	c.Data(200, contentType, bytes)
}

func NewImage(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		error := handlers.ErrMissingData

		c.JSON(error.Status, gin.H{
			"success": false,
			"message": strings.Replace(error.Error, "{}", "file", 1),
		})
		return
	}

	domain := c.PostForm("domain")

	if domain == "" {
		error := handlers.ErrMissingData

		c.JSON(error.Status, gin.H{
			"success": false,
			"message": strings.Replace(error.Error, "{}", "domain", 1),
		})
		return
	}

	if !utils.CheckIfElementExists(config.Config.Domains, domain) {
		domains := strings.Join(config.Config.Domains, ", ")
		error := handlers.ErrMissingData

		c.JSON(error.Status, gin.H{
			"success": false,
			"message": strings.Replace(strings.Replace(error.Error, "{}", "domain", 1), "{}", domains, 1),
		})
		return
	}

	if file.Size > config.Config.MaxFileSize {
		error := handlers.ErrFileTooBig

		c.JSON(error.Status, gin.H{
			"success": false,
			"message": strings.Replace(error.Error, "{}", units.BinarySuffix(float64(config.Config.MaxFileSize)), 1),
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
		error := handlers.ErrInvalid

		c.JSON(error.Status, gin.H{
			"success": false,
			"message": strings.Replace(strings.Replace(error.Error, "{}", "file type", 1), "{}", "jpg, png and gif", 1),
		})
		return
	}

	// NOTE: These are for me, dont mind them
	// 🍖🍗🍘🍙🍚🍛🍜🍝🍞🍟🍠🍡🍢🍣🍤🍥🍦🍧🍨🍩🍪🍫🍬🍭🍮🍯🍰🎊🎋🎌🎍🎎🎏🎑🎒🎓🎠🎡🎢🎣🎤🎥🎦🎧🎨🎩🎪🎬🏀🏁🐀🐁🐂🐃🐄🐅🐆🐇🐈🐉🐊🐋🐌
	// 255 emojis i can use
	// ⌚⌛⏩⏪⏫⏬⏰⏳◽◾☔☕♈♉♊♋♌♍♎♏♐♑♒♓♿⚓⚡⚪⚫⚽⚾⛄⛅⛎⛔⛪⛲⛳⛵⛺⛽✅✊✋✨❌❎❓❔❕❗➕➖➗➰➿⬛⬜⭐⭕🀄🃏🆎🆑🆒🆓🆔🆕🆖🆗🆘🆙🆚🈁🈚🈯🈲🈳🈴🈵🈶🈸🈹🈺🉐🉑🌀🌁🌂🌃🌄🌅🌆🌇🌈🌉🌊🌋🌌🌍🌎🌏🌐🌑🌒🌓🌔🌕🌖🌗🌘🌙🌚🌛🌜🌝🌞🌟🌠🌭🌮🌯🌰🌱🌲🌳🌴🌵🌷🌸🌹🌺🌻🌼🌽🌾🌿🍀🍁🍂🍃🍄🍅🍆🍇🍈🍉🍊🍋🍌🍍🍎🍏🍐🍑🍒🍓🍔🍕🍖🍗🍘🍙🍚🍛🍜🍝🍞🍟🍠🍡🍢🍣🍤🍥🍦🍧🍨🍩🍪🍫🍬🍭🍮🍯🍰🍱🍲🍳🍴🍵🍶🍷🍸🍹🍺🍻🍼🍾🍿🎀🎁🎂🎃🎄🎅🎆🎇🎈🎉🎊🎋🎌🎍🎎🎏🎐🎑🎒🎓🎠🎡🎢🎣🎤🎥🎦🎧🎨🎩🎪🎫🎬🎭🎮🎯🎰🎱🎲🎳🎴🎵🎶🎷🎸🎹🎺🎻🎼🎽🎾🎿🏀🏁🏂🏃🏄🏅🏆🏇🏈🏉🏊🏏🏐🏑🏒🏓🏠🏡🏢🏣🏤🏥🏦🏧🏨🏩🏪🏫🏬🏭🏮🏯🏰🏴🏸🏹🏺🏻🏼🏽🏾🏿🐀🐁🐂🐃🐄🐅🐆🐇🐈🐉🐊🐋🐌🐍🐎🐏🐐🐑🐒🐓🐔🐕🐖🐗🐘🐙🐚🐛🐜🐝🐞🐟🐠🐡🐢🐣🐤🐥🐦🐧🐨🐩🐪🐫🐬🐭🐮🐯🐰🐱🐲🐳🐴🐵🐶🐷🐸🐹🐺🐻🐼🐽🐾👀👂👃👄👅👆👇👈👉👊👋👌👍👎👏👐👑👒👓👔👕👖👗👘👙👚👛👜👝👞👟👠👡👢👣👤👥👦👧👨👩👪👫👬👭👮👯👰👱👲👳👴👵👶👷👸👹👺👻👼👽👾👿💀💁💂💃💄💅💆💇💈💉💊💋💌💍💎💏💐💑💒💓💔💕💖💗💘💙💚💛💜💝💞💟💠💡💢💣💤💥💦💧💨💩💪💫💬💭💮💯💰💱💲💳💴💵💶💷💸💹💺💻💼💽💾💿📀📁📂📃📄📅📆📇📈📉📊📋📌📍📎📏📐📑📒📓📔📕📖📗📘📙📚📛📜📝📞📟📠📡📢📣📤📥📦📧📨📩📪📫📬📭📮📯📰📱📲📳📴📵📶📷📸📹📺📻📼📿🔀🔁🔂🔃🔄🔅🔆🔇🔈🔉🔊🔋🔌🔍🔎🔏🔐🔑🔒🔓🔔🔕🔖🔗🔘🔙🔚🔛🔜🔝🔞🔟🔠🔡🔢🔣🔤🔥🔦🔧🔨🔩🔪🔫🔬🔭🔮🔯🔰🔱🔲🔳🔴🔵🔶🔷🔸🔹🔺🔻🔼🔽🕋🕌🕍🕎🕐🕑🕒🕓🕔🕕🕖🕗🕘🕙🕚🕛🕜🕝🕞🕟🕠🕡🕢🕣🕤🕥🕦🕧🕺🖕🖖🖤🗻🗼🗽🗾🗿😀😁😂😃😄😅😆😇😈😉😊😋😌😍😎😏😐😑😒😓😔😕😖😗😘😙😚😛😜😝😞😟😠😡😢😣😤😥😦😧😨😩😪😫😬😭😮😯😰😱😲😳😴😵😶😷😸😹😺😻😼😽😾😿🙀🙁🙂🙃🙄🙅🙆🙇🙈🙉🙊🙋🙌🙍🙎🙏🚀🚁🚂🚃🚄🚅🚆🚇🚈🚉🚊🚋🚌🚍🚎🚏🚐🚑🚒🚓🚔🚕🚖🚗🚘🚙🚚🚛🚜🚝🚞🚟🚠🚡🚢🚣🚤🚥🚦🚧🚨🚩🚪🚫🚬🚭🚮🚯🚰🚱🚲🚳🚴🚵🚶🚷🚸🚹🚺🚻🚼🚽🚾🚿🛀🛁🛂🛃🛄🛅🛌🛐🛑🛒🛕🛖🛗🛫🛬🛴🛵🛶🛷🛸🛹🛺🛻🛼🟠🟡🟢🟣🟤🟥🟦🟧🟨🟩🟪🟫🤌🤍🤎🤏🤐🤑🤒🤓🤔🤕🤖🤗🤘🤙🤚🤛🤜🤝🤞🤟🤠🤡🤢🤣🤤🤥🤦🤧🤨🤩🤪🤫🤬🤭🤮🤯🤰🤱🤲🤳🤴🤵🤶🤷🤸🤹🤺🤼🤽🤾🤿🥀🥁🥂🥃🥄🥅🥇🥈🥉🥊🥋🥌🥍🥎🥏🥐🥑🥒🥓🥔🥕🥖🥗🥘🥙🥚🥛🥜🥝🥞🥟🥠🥡🥢🥣🥤🥥🥦🥧🥨🥩🥪🥫🥬🥭🥮🥯🥰🥱🥲🥳🥴🥵🥶🥷🥸🥺🥻🥼🥽🥾🥿🦀🦁🦂🦃🦄🦅🦆🦇🦈🦉🦊🦋🦌🦍🦎🦏🦐🦑🦒🦓🦔🦕🦖🦗🦘🦙🦚🦛🦜🦝🦞🦟🦠🦡🦢🦣🦤🦥🦦🦧🦨🦩🦪🦫🦬🦭🦮🦯🦰🦱🦲🦳🦴🦵🦶🦷🦸🦹🦺🦻🦼🦽🦾🦿🧀🧁🧂🧃🧄🧅🧆🧇🧈🧉🧊🧋🧍🧎🧏🧐🧑🧒🧓🧔🧕🧖🧗🧘🧙🧚🧛🧜🧝🧞🧟🧠🧡🧢🧣🧤🧥🧦🧧🧨🧩🧪🧫🧬🧭🧮🧯🧰🧱🧲🧳🧴🧵🧶🧷🧸🧹🧺🧻🧼🧽🧾🧿🩰🩱🩲🩳🩴🩸🩹🩺🪀🪁🪂🪃🪄🪅🪆🪐🪑🪒🪓🪔🪕🪖🪗🪘🪙🪚🪛🪜🪝🪞🪟🪠🪡🪢🪣🪤🪥🪦🪧🪨🪰🪱🪲🪳🪴🪵🪶🫀🫁🫂🫐🫑🫒🫓🫔🫕🫖©️®️‼️⁉️™️ℹ️↔️↕️↖️↗️↘️↙️↩️↪️⌨️⏏️⏭️⏮️⏯️⏱️⏲️⏸️⏹️⏺️Ⓜ️▪️▫️▶️◀️◻️◼️☀️☁️☂️☃️☄️☎️☑️☘️☝️☠️☢️☣️☦️☪️☮️☯️☸️☹️☺️♀️♂️♟️♠️♣️♥️♦️♨️♻️♾️⚒️⚔️⚕️⚖️⚗️⚙️⚛️⚜️⚠️⚧️⚰️⚱️⛈️⛏️⛑️⛓️⛩️⛰️⛱️⛴️⛷️⛸️⛹️✂️✈️✉️✌️✍️✏️✒️✔️✖️✝️✡️✳️✴️❄️❇️❣️❤️➡️⤴️⤵️⬅️⬆️⬇️〰️〽️㊗️㊙️🅰️🅱️🅾️🅿️🈂️🈷️🌡️🌤️🌥️🌦️🌧️🌨️🌩️🌪️🌫️🌬️🌶️🍽️🎖️🎗️🎙️🎚️🎛️🎞️🎟️🏋️🏌️🏍️🏎️🏔️🏕️🏖️🏗️🏘️🏙️🏚️🏛️🏜️🏝️🏞️🏟️🏳️🏵️🏷️🐿️👁️📽️🕉️🕊️🕯️🕰️🕳️🕴️🕵️🕶️🕷️🕸️🕹️🖇️🖊️🖋️🖌️🖍️🖐️🖥️🖨️🖱️🖲️🖼️🗂️🗃️🗄️🗑️🗒️🗓️🗜️🗝️🗞️🗡️🗣️🗨️🗯️🗳️🗺️🛋️🛍️🛎️🛏️🛠️🛡️🛢️🛣️🛤️🛥️🛩️🛰️🛳️#️⃣*️⃣0️⃣1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣
	id, err := gonanoid.New(config.Config.IdLength)

	if err != nil {
		error := handlers.ErrUnknownErrorOccurred

		c.JSON(error.Status, gin.H{
			"success": false,
			"message": error.Error,
		})
		return
	}

	name := fmt.Sprintf("%s.%s", id, extension)

	src, err := file.Open()
	if err != nil {
		error := handlers.ErrUnknownErrorOccurred

		c.JSON(error.Status, gin.H{
			"success": false,
			"message": error.Error,
		})
		return
	}
	defer src.Close()

	err = storage.Storage.Put(src, name)

	if err != nil {
		error := handlers.ErrUnknownErrorOccurred

		c.JSON(error.Status, gin.H{
			"success": false,
			"message": error.Error,
		})
		return
	}

	url := fmt.Sprintf("https://%s/%s", domain, name)

	c.JSON(200, gin.H{
		"success": true,
		"message": url,
	})
}
