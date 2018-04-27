package froniusCurrentPowerflow

type CurrentPowerflow struct {
	Site      Site
	Inverters []Inverter
	Ohmpilots []Ohmpilot
}

func (currentPowerflow CurrentPowerflow) Persist() error {
	site, err := currentPowerflow.Site.Persist()
	if err != nil {
		return err
	}

	for _, inverter := range currentPowerflow.Inverters {
		if err := inverter.Persist(site.Id); err != nil {
			return err
		}
	}

	for _, ohmpilot := range currentPowerflow.Ohmpilots {
		if err := ohmpilot.Persist(site.Id); err != nil {
			return err
		}
	}

	return nil
}