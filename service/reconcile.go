package service

// reconcileLocalAgainstRemote treats remote as the source of truth for a set of
// local records keyed by ID. Locals absent from remote are pruned via onMissing;
// locals present remotely are synced via onFound.
func reconcileLocalAgainstRemote[T any, R any](
	locals []T,
	idOf func(T) string,
	remote map[string]R,
	onMissing func(T) error,
	onFound func(T, R) error,
) error {
	for _, local := range locals {
		rem, ok := remote[idOf(local)]
		if !ok {
			if err := onMissing(local); err != nil {
				return err
			}
			continue
		}
		if err := onFound(local, rem); err != nil {
			return err
		}
	}
	return nil
}
