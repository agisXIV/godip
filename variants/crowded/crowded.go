package crowded

import (
	"github.com/zond/godip"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/start"
	"github.com/zond/godip/variants/common"
)

var FranceAustriaVariant = common.Variant{
	Name: "Classic: Crowded",
	Graph: func() godip.Graph {
		okNations := map[godip.Nation]bool{
			godip.England: true,
      godip.France:  true,
      godip.Iberia: true,
      godip.Benelux: true,
      godip.Scandinavia: true,
      godip.Germany: true,
      godip.Italy: true,
      godip.Russia: true,
      godip.Balkans: true,
			godip.Austria: true,
      godip.Turkey: true,
			godip.Neutral: true,
		}
		neutral := godip.Neutral
		result := start.Graph()
		for _, node := range result.Nodes {
			if node.SC != nil && !okNations[*node.SC] {
				node.SC = &neutral
			}
		}
		return result
	},
	Start: func() (result *state.State, err error) {
		if result, err = classical.Start(); err != nil {
			return
		}
		if err = result.SetUnits(map[godip.Province]godip.Unit{
			"bre": godip.Unit{godip.Fleet, godip.France},
			"par": godip.Unit{godip.Army, godip.France},
			"mar": godip.Unit{godip.Army, godip.France},
			"tri": godip.Unit{godip.Fleet, godip.Austria},
			"vie": godip.Unit{godip.Army, godip.Austria},
			"bud": godip.Unit{godip.Army, godip.Austria},
      "spa": godip.Unit{godip.Army, godip.Iberia},
      "por": godip.Unit{godip.Fleet, godip.Iberia},
      "tun": godip.Unit{godip.Fleet, godip.Iberia},
      "lvp": godip.Unit{godip.Fleet, godip.England},
      "edi": godip.Unit{godip.Fleet, godip.England},
      "liv": godip.Unit{godip.Army, godip.England},
      "bel": godip.Unot{godip.Fleet, godip.Benelux},
      "hol": godip.Unit{godip.Army, godip.Benelux},
      "ruh": godip.Unit{godip.Army, godip.Benelux},
      "ber": godip.Unit{godip.Army, godip.Germany},
      "kie": godip.Unit{godip.Fleet, godip.Germany},
      "mun": godip.Unit{godip.Army, godip.Germany},
      "nwy": godip.Unit{godip.Fleet, godip.Scandinavia},
      "swe"  godip.Unit{godip.Army, godip.Scandinavia},
      "den": godip.Unit{godip.Fleet, godip.Scandinavia},
      "stp/sc": godip.Unit{godip.Fleet, godip.Russia},
      "mos": godip.Unit{godip.Army, godip.Russia},
      "war": godip.Unit{godip.Army, godip.Russia},
      "sev": godip.Unit{godip.Fleer, godip.Russia},
      "rum": godip.Unit{godip.Fleet, godip.Balkan},
      "gre": godip.Unit{godip.Fleet, godip.Balkan},
      "ser": godip.Unit{godip.Army, godip.Balkan},
      "bul": godip.Unit{godip.Army, godip.Balkan},
      "ank": godip.Unit{godip.Army, godip.Turkey},
      "con": godip.Unit{godip.Army, godip.Turkey},
      "smy": godip.Unit{godip.Fleet, godip.Turkey},
      "rom": godip.Unit{godip.Army, godip.Italy},
      "ven": godip.Unit{godip.Army, godip.Italy},
      "nap": godip.Unit{godip.Fleet, godip.Italy},
      
		}); err != nil {
			return
		}
		result.SetSupplyCenters(map[godip.Province]godip.Nation{
			"bre": godip.France,
			"par": godip.France,
			"mar": godip.France,
			"tri": godip.Austria,
			"vie": godip.Austria,
			"bud": godip.Austria,
      "spa": godip.Iberia,
      "por": godip.Iberia,
      "tun": godip.Iberia,
      "lon": godip.England,
      "lvp": godip.England,
      "edi": godip.England,
      "nwy": godip.Scandinavia,
      "swe": godip.Scandinavia,
      "den": godip.Scandinavia,
      "ber": godip.Germany,
      "kie": godip.Germany,
      "mun": godip.Germany,
      "ruh": godip.Benelux,
      "bel": godip.Benelux,
      "hol": godip.Benelux,
      "stp": godip.Russia,
      "mos": godip.Russia,
      "war": godip.Russia,
      "sev": godip.Russia,
      "rum": godip.Balkan,
      "gre": godip.Balkan,
      "ser": godip.Balkan,
      "bul": godip.Balkan,
      "ank": godip.Turkey,
      "smy": godip.Turkey,
      "con": godip.Turkey,
      "rom": godip.Italy,
      "nap": godip.Italy,
      "ven": godip.Italy,
      
		})
		return
	},
	Blank:             classical.Blank,
	Phase:             classical.NewPhase,
	Parser:            classical.Parser,
	Nations:           []godip.Nation{godip.Austria, godip.France, godip.England, godip.Iberia, godip.Germany, godip.Benelux, godip.Russia, godip.Italy, godip.Scandinavia, godip.Balkan, godip.Turkey},
	PhaseTypes:        classical.PhaseTypes,
	Seasons:           classical.Seasons,
	UnitTypes:         classical.UnitTypes,
	SoloWinner:        common.SCCountWinner(18),
	ProvinceLongNames: classical.ClassicalVariant.ProvinceLongNames,
	SVGMap: func() ([]byte, error) {
		return classical.Asset("svg/map.svg")
	},
	SVGVersion:  "1",
	SVGUnits:    classical.SVGUnits,
	CreatedBy:   "",
	Version:     "",
	Description: "The classical map with 4 extra nations packed in so as to have no neutral scs. Also Ruhr is a supply centre now.",
	SoloSCCount: func(*state.State) int { return 18 },
	Rules: `The first to 18 supply centers is the winner. 
	Normal dip rules apply`,
}
