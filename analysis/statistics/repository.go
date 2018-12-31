package statistics

type repository interface {
	Write()

	Get(event string)

	Close()
}
