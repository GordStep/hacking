package status

import "math/rand"

// Зелёный
var greenStatuses = []string{
	"   STARTING    ",
	"   WORKING     ",
	"   RUNNING     ",
	"   SCALING     ",
	"   EXPLOITING  ",
	"   INJECTED    ",
	"   ROOTED      ",
	"   OWNED       ",
	"   DEPLOYED    ",
	"   PROVISIONED ",
	"   CONNECTED   ",
	"   TRAINING    ",
	"   CRAFTING    ",
	"   OK          ",
	"   SUCCESS     ",
	"   ONLINE      ",
	"   ACTIVE      ",
	"   HEALTHY     ",
	"   OPTIMAL     ",
}

// Жёлтый
var yellowStatuses = []string{
	"   PROCESSING  ",
	"   COMPILING   ",
	"   INDEXING    ",
	"   RESOLVING   ",
	"   ENUMERATING ",
	"   PROBING     ",
	"   SCANNING    ",
	"   FINGERPRINT ",
	"   BUILDING    ",
	"   TESTING     ",
	"   PUSHING     ",
	"   MERGING     ",
	"   ROUTING     ",
	"   PREDICTING  ",
	"   THINKING    ",
	"   PLEASE WAIT ",
	"   PENDING     ",
	"   QUEUEING    ",
	"   BUFFERING   ",
	"   BAKING      ",
}

// Белый
var whiteStatuses = []string{
	"   LOADING     ",
	"   FETCHING    ",
	"   DOWNLOADING ",
	"   READING     ",
	"   DECRYPTING  ",
	"   DUMPING     ",
	"   SNIFFING    ",
	"   PULLING     ",
	"   CACHING     ",
	"   LOGGING     ",
	"   MONITORING  ",
	"   GENERATING  ",
	"   LOOTING     ",
	"   SYNCING     ",
	"   CLONING     ",
	"   ARCHIVING   ",
	"   STREAMING   ",
	"   PARSING     ",
	"   UNPACKING   ",
}

// Красный
var redStatuses = []string{
	"   STOPPING    ",
	"   FAILING     ",
	"   CRASHING    ",
	"   BREACHING   ",
	"   DETECTED    ",
	"   TRACED      ",
	"   BURNED      ",
	"   LOCKDOWN    ",
	"   ROLLBACK    ",
	"   OUTAGE      ",
	"   DEGRADED    ",
	"   THROTTLED   ",
	"   DISCONNECT  ",
	"   OVERFITTING ",
	"   FAIL        ",
	"   ABORTED     ",
	"   CORRUPTED   ",
	"   TIMEOUT     ",
	"   BSOD        ",
}

// Get возвращает случайный статус по цвету
func Get(color string) string {
	switch color {
	case "green":
		return greenStatuses[rand.Intn(len(greenStatuses))]
	case "yellow":
		return yellowStatuses[rand.Intn(len(yellowStatuses))]
	case "white":
		return whiteStatuses[rand.Intn(len(whiteStatuses))]
	case "red":
		return redStatuses[rand.Intn(len(redStatuses))]
	default:
		return greenStatuses[rand.Intn(len(greenStatuses))]
	}
}

// GetByNum возвращает случайный статус по числовому коду col
func GetByNum(col int) string {
	color := GetColorByNum(col)
	return Get(color)
}

// GetColorByNum возвращает цвет по числовому коду col
func GetColorByNum(col int) string {
	switch col {
	case 1:
		return "white"
	case 4:
		return "red"
	case 6:
		return "yellow"
	default:
		return "green"
	}
}
