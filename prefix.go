package id

type prefix string

func (p prefix) isValid(acceptedKeys ...string) bool {
	containsKey := false
	for _, v := range prefixes {
		if v == p {
			containsKey = true
			break
		}
	}

	if !containsKey {
		return false
	}

	if len(acceptedKeys) == 0 {
		return containsKey
	}

	for _, key := range acceptedKeys {
		if prefixes[key] == p {
			return true
		}
	}
	return false
}

const (
	accountPrefix         prefix = "acct"
	alternativePrefix     prefix = "alt"
	categoryPrefix        prefix = "cat"
	commentPrefix         prefix = "cmt"
	authPrefix            prefix = "auth"
	contactPrefix         prefix = "cont"
	filePrefix            prefix = "file"
	productPrefix         prefix = "prod"
	projectPrefix         prefix = "proj"
	requestPrefix         prefix = "req"
	scheduleItemPrefix    prefix = "schi"
	schedulePrefix        prefix = "sch"
	similarityScorePrefix prefix = "scr"
	supplierPrefix        prefix = "sup"
	supplierContactPrefix prefix = "supc"
	updateMessagePrefix   prefix = "msg"
	userPrefix            prefix = "usr"
)

var prefixes = map[string]prefix{
	"account":          accountPrefix,
	"alternative":      alternativePrefix,
	"auth":             authPrefix,
	"category":         categoryPrefix,
	"comment":          commentPrefix,
	"contact":          contactPrefix,
	"file":             filePrefix,
	"product":          productPrefix,
	"project":          projectPrefix,
	"request":          requestPrefix,
	"schedule":         schedulePrefix,
	"schedule_item":    scheduleItemPrefix,
	"similarity_score": similarityScorePrefix,
	"supplier":         supplierPrefix,
	"supplier_contact": supplierContactPrefix,
	"update_message":   updateMessagePrefix,
	"user":             userPrefix,
}
