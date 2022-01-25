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
		error, status := handlers.ErrMissingData("id")

		c.JSON(status, gin.H{
			"success": false,
			"message": error,
		})
		return
	}

	reader, err := storage.Storage.Get(data.Id)

	if err != nil {
		error, status := handlers.ErrResourceNotFound()

		c.JSON(status, gin.H{
			"success": false,
			"message": error,
		})
		return
	}

	bytes := utils.IoReaderToByteSlice(reader)

	bytes512 := make([]byte, 512)
	copy(bytes512, bytes)

	contentType := http.DetectContentType(bytes512)

	if err != nil {
		error, status := handlers.ErrUnknownErrorOccurred()

		c.JSON(status, gin.H{
			"success": false,
			"message": error,
		})
		return
	}

	c.Data(200, contentType, bytes)
}

func NewImage(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		error, status := handlers.ErrMissingData("file")

		c.JSON(status, gin.H{
			"success": false,
			"message": error,
		})
		return
	}

	domain := c.PostForm("domain")

	if domain == "" {
		error, status := handlers.ErrMissingData("domain")

		c.JSON(status, gin.H{
			"success": false,
			"message": error,
		})
		return
	}

	if !utils.CheckIfElementExists(config.Config.Domains, domain) {
		domains := strings.Join(config.Config.Domains, ", ")
		error, status := handlers.ErrInvalidData("domain", domains)

		c.JSON(status, gin.H{
			"success": false,
			"message": error,
		})
		return
	}

	if file.Size > config.Config.MaxFileSize {
		error, status := handlers.ErrFileTooBig(units.BinarySuffix(float64(config.Config.MaxFileSize)))

		c.JSON(status, gin.H{
			"success": false,
			"message": error,
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
		error, status := handlers.ErrInvalidData("file type", "jpg, png and gif")

		c.JSON(status, gin.H{
			"success": false,
			"message": error,
		})
		return
	}

	// NOTE: These are for me, dont mind them
	// 🍖🍗🍘🍙🍚🍛🍜🍝🍞🍟🍠🍡🍢🍣🍤🍥🍦🍧🍨🍩🍪🍫🍬🍭🍮🍯🍰🎊🎋🎌🎍🎎🎏🎑🎒🎓🎠🎡🎢🎣🎤🎥🎦🎧🎨🎩🎪🎬🏀🏁🐀🐁🐂🐃🐄🐅🐆🐇🐈🐉🐊🐋🐌
	// 255 emojis i can use
	// ⌚⌛⏩⏪⏫⏬⏰⏳◽◾☔☕♈♉♊♋♌♍♎♏♐♑♒♓♿⚓⚡⚪⚫⚽⚾⛄⛅⛎⛔⛪⛲⛳⛵⛺⛽✅✊✋✨❌❎❓❔❕❗➕➖➗➰➿⬛⬜⭐⭕🀄🃏🆎🆑🆒🆓🆔🆕🆖🆗🆘🆙🆚🈁🈚🈯🈲🈳🈴🈵🈶🈸🈹🈺🉐🉑🌀🌁🌂🌃🌄🌅🌆🌇🌈🌉🌊🌋🌌🌍🌎🌏🌐🌑🌒🌓🌔🌕🌖🌗🌘🌙🌚🌛🌜🌝🌞🌟🌠🌭🌮🌯🌰🌱🌲🌳🌴🌵🌷🌸🌹🌺🌻🌼🌽🌾🌿🍀🍁🍂🍃🍄🍅🍆🍇🍈🍉🍊🍋🍌🍍🍎🍏🍐🍑🍒🍓🍔🍕🍖🍗🍘🍙🍚🍛🍜🍝🍞🍟🍠🍡🍢🍣🍤🍥🍦🍧🍨🍩🍪🍫🍬🍭🍮🍯🍰🍱🍲🍳🍴🍵🍶🍷🍸🍹🍺🍻🍼🍾🍿🎀🎁🎂🎃🎄🎅🎆🎇🎈🎉🎊🎋🎌🎍🎎🎏🎐🎑🎒🎓🎠🎡🎢🎣🎤🎥🎦🎧🎨🎩🎪🎫🎬🎭🎮🎯🎰🎱🎲🎳🎴🎵🎶🎷🎸🎹🎺🎻🎼🎽🎾🎿🏀🏁🏂🏃🏄🏅🏆🏇🏈🏉🏊🏏🏐🏑🏒🏓🏠🏡🏢🏣🏤🏥🏦🏧🏨🏩🏪🏫🏬🏭🏮🏯🏰🏴🏸🏹🏺🏻🏼🏽🏾🏿🐀🐁🐂🐃🐄🐅🐆🐇🐈🐉🐊🐋🐌🐍🐎🐏🐐🐑🐒🐓🐔🐕🐖🐗🐘🐙🐚🐛🐜🐝🐞🐟🐠🐡🐢🐣🐤🐥🐦🐧🐨🐩🐪🐫🐬🐭🐮🐯🐰🐱🐲🐳🐴🐵🐶🐷🐸🐹🐺🐻🐼🐽🐾👀👂👃👄👅👆👇👈👉👊👋👌👍👎👏👐👑👒👓👔👕👖👗👘👙👚👛👜👝👞👟👠👡👢👣👤👥👦👧👨👩👪👫👬👭👮👯👰👱👲👳👴👵👶👷👸👹👺👻👼👽👾👿💀💁💂💃💄💅💆💇💈💉💊💋💌💍💎💏💐💑💒💓💔💕💖💗💘💙💚💛💜💝💞💟💠💡💢💣💤💥💦💧💨💩💪💫💬💭💮💯💰💱💲💳💴💵💶💷💸💹💺💻💼💽💾💿📀📁📂📃📄📅📆📇📈📉📊📋📌📍📎📏📐📑📒📓📔📕📖📗📘📙📚📛📜📝📞📟📠📡📢📣📤📥📦📧📨📩📪📫📬📭📮📯📰📱📲📳📴📵📶📷📸📹📺📻📼📿🔀🔁🔂🔃🔄🔅🔆🔇🔈🔉🔊🔋🔌🔍🔎🔏🔐🔑🔒🔓🔔🔕🔖🔗🔘🔙🔚🔛🔜🔝🔞🔟🔠🔡🔢🔣🔤🔥🔦🔧🔨🔩🔪🔫🔬🔭🔮🔯🔰🔱🔲🔳🔴🔵🔶🔷🔸🔹🔺🔻🔼🔽🕋🕌🕍🕎🕐🕑🕒🕓🕔🕕🕖🕗🕘🕙🕚🕛🕜🕝🕞🕟🕠🕡🕢🕣🕤🕥🕦🕧🕺🖕🖖🖤🗻🗼🗽🗾🗿😀😁😂😃😄😅😆😇😈😉😊😋😌😍😎😏😐😑😒😓😔😕😖😗😘😙😚😛😜😝😞😟😠😡😢😣😤😥😦😧😨😩😪😫😬😭😮😯😰😱😲😳😴😵😶😷😸😹😺😻😼😽😾😿🙀🙁🙂🙃🙄🙅🙆🙇🙈🙉🙊🙋🙌🙍🙎🙏🚀🚁🚂🚃🚄🚅🚆🚇🚈🚉🚊🚋🚌🚍🚎🚏🚐🚑🚒🚓🚔🚕🚖🚗🚘🚙🚚🚛🚜🚝🚞🚟🚠🚡🚢🚣🚤🚥🚦🚧🚨🚩🚪🚫🚬🚭🚮🚯🚰🚱🚲🚳🚴🚵🚶🚷🚸🚹🚺🚻🚼🚽🚾🚿🛀🛁🛂🛃🛄🛅🛌🛐🛑🛒🛕🛖🛗🛫🛬🛴🛵🛶🛷🛸🛹🛺🛻🛼🟠🟡🟢🟣🟤🟥🟦🟧🟨🟩🟪🟫🤌🤍🤎🤏🤐🤑🤒🤓🤔🤕🤖🤗🤘🤙🤚🤛🤜🤝🤞🤟🤠🤡🤢🤣🤤🤥🤦🤧🤨🤩🤪🤫🤬🤭🤮🤯🤰🤱🤲🤳🤴🤵🤶🤷🤸🤹🤺🤼🤽🤾🤿🥀🥁🥂🥃🥄🥅🥇🥈🥉🥊🥋🥌🥍🥎🥏🥐🥑🥒🥓🥔🥕🥖🥗🥘🥙🥚🥛🥜🥝🥞🥟🥠🥡🥢🥣🥤🥥🥦🥧🥨🥩🥪🥫🥬🥭🥮🥯🥰🥱🥲🥳🥴🥵🥶🥷🥸🥺🥻🥼🥽🥾🥿🦀🦁🦂🦃🦄🦅🦆🦇🦈🦉🦊🦋🦌🦍🦎🦏🦐🦑🦒🦓🦔🦕🦖🦗🦘🦙🦚🦛🦜🦝🦞🦟🦠🦡🦢🦣🦤🦥🦦🦧🦨🦩🦪🦫🦬🦭🦮🦯🦰🦱🦲🦳🦴🦵🦶🦷🦸🦹🦺🦻🦼🦽🦾🦿🧀🧁🧂🧃🧄🧅🧆🧇🧈🧉🧊🧋🧍🧎🧏🧐🧑🧒🧓🧔🧕🧖🧗🧘🧙🧚🧛🧜🧝🧞🧟🧠🧡🧢🧣🧤🧥🧦🧧🧨🧩🧪🧫🧬🧭🧮🧯🧰🧱🧲🧳🧴🧵🧶🧷🧸🧹🧺🧻🧼🧽🧾🧿🩰🩱🩲🩳🩴🩸🩹🩺🪀🪁🪂🪃🪄🪅🪆🪐🪑🪒🪓🪔🪕🪖🪗🪘🪙🪚🪛🪜🪝🪞🪟🪠🪡🪢🪣🪤🪥🪦🪧🪨🪰🪱🪲🪳🪴🪵🪶🫀🫁🫂🫐🫑🫒🫓🫔🫕🫖©️®️‼️⁉️™️ℹ️↔️↕️↖️↗️↘️↙️↩️↪️⌨️⏏️⏭️⏮️⏯️⏱️⏲️⏸️⏹️⏺️Ⓜ️▪️▫️▶️◀️◻️◼️☀️☁️☂️☃️☄️☎️☑️☘️☝️☠️☢️☣️☦️☪️☮️☯️☸️☹️☺️♀️♂️♟️♠️♣️♥️♦️♨️♻️♾️⚒️⚔️⚕️⚖️⚗️⚙️⚛️⚜️⚠️⚧️⚰️⚱️⛈️⛏️⛑️⛓️⛩️⛰️⛱️⛴️⛷️⛸️⛹️✂️✈️✉️✌️✍️✏️✒️✔️✖️✝️✡️✳️✴️❄️❇️❣️❤️➡️⤴️⤵️⬅️⬆️⬇️〰️〽️㊗️㊙️🅰️🅱️🅾️🅿️🈂️🈷️🌡️🌤️🌥️🌦️🌧️🌨️🌩️🌪️🌫️🌬️🌶️🍽️🎖️🎗️🎙️🎚️🎛️🎞️🎟️🏋️🏌️🏍️🏎️🏔️🏕️🏖️🏗️🏘️🏙️🏚️🏛️🏜️🏝️🏞️🏟️🏳️🏵️🏷️🐿️👁️📽️🕉️🕊️🕯️🕰️🕳️🕴️🕵️🕶️🕷️🕸️🕹️🖇️🖊️🖋️🖌️🖍️🖐️🖥️🖨️🖱️🖲️🖼️🗂️🗃️🗄️🗑️🗒️🗓️🗜️🗝️🗞️🗡️🗣️🗨️🗯️🗳️🗺️🛋️🛍️🛎️🛏️🛠️🛡️🛢️🛣️🛤️🛥️🛩️🛰️🛳️#️⃣*️⃣0️⃣1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣
	id, err := gonanoid.New(config.Config.IdLength)

	if err != nil {
		error, status := handlers.ErrUnknownErrorOccurred()

		c.JSON(status, gin.H{
			"success": false,
			"message": error,
		})
		return
	}

	name := fmt.Sprintf("%s.%s", id, extension)

	src, err := file.Open()
	if err != nil {
		error, status := handlers.ErrUnknownErrorOccurred()

		c.JSON(status, gin.H{
			"success": false,
			"message": error,
		})
		return
	}
	defer src.Close()

	err = storage.Storage.Put(src, name)

	if err != nil {
		error, status := handlers.ErrUnknownErrorOccurred()

		c.JSON(status, gin.H{
			"success": false,
			"message": error,
		})
		return
	}

	url := fmt.Sprintf("https://%s/%s", domain, name)

	c.JSON(200, gin.H{
		"success": true,
		"message": url,
	})
}
