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
		inverter.Persist(site.Id)
	}

	for _, ohmpilot := range currentPowerflow.Ohmpilots {
		ohmpilot.Persist(site.Id)
	}

	return nil
}