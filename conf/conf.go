package conf

var cfg *RobotConf

func init() {
	cfg = new(RobotConf)
	cfg.CfgMatchidRobots = make(map[string]*CfgMatchidRobot)
}

type RobotConf struct {
	CfgMatchidRobots map[string]*CfgMatchidRobot
}

type CfgMatchidRobot struct {
	Total  int
	Status int
}

func GetCfgMatchidRobot() map[string]*CfgMatchidRobot {
	return cfg.CfgMatchidRobots
}
