package main

import (
	"fmt"
	"regexp"
)

const text = `
my email list:
- WisdomFusion@Gmail.com
- huzhifei@Live.com
- gavin2u@sina.com.cn

and QQ is 664009005
`

func main() {
	reMail := regexp.MustCompile(`\w+@\w+\.\w+`)
	reQQ := regexp.MustCompile(`\d{5,}`)

	mails := reMail.FindAllString(text, -1)
	QQ := reQQ.FindString(text)

	for _, mail := range mails {
		fmt.Println(mail)
	}

	fmt.Println(QQ)

	fmt.Println("============")

	reSubmatch := regexp.MustCompile(`(\w+)@(\w.+\.\w+)`)

	reSubmatches := reSubmatch.FindAllStringSubmatch(text, -1)

	for _, m := range reSubmatches {
		fmt.Println(m)

		for _, s := range m {
			fmt.Println(s)
		}
	}
}
