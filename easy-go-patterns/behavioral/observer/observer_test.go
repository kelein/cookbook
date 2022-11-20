package observer

import (
	"log"
	"testing"
)

func TestInformerNotify(t *testing.T) {
	t.Run("A", func(t *testing.T) {
		hero1 := &Hero{"HuanRon", "WUDAN"}
		hero2 := &Hero{"Hqigong", "WUDAN"}
		hero3 := &Hero{"QiaoFeng", "WUDAN"}
		hero4 := &Hero{"ZhanWuji", "ERMEI"}
		hero5 := &Hero{"BaXiXiao", "ERMEI"}
		hero6 := &Hero{"LionKing", "ERMEI"}

		informer := &Informer{}
		informer.Attach(hero1)
		informer.Attach(hero2)
		informer.Attach(hero3)
		informer.Attach(hero4)
		informer.Attach(hero5)
		informer.Attach(hero6)
		log.Printf("informer with slience, count: %v", informer.count)
		hero1.Fight(hero5, informer)
	})
}
