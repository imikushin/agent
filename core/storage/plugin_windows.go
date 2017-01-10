package storage

// CallVolumeDriverAttach is not supported on windows
func CallVolumeDriverAttach(volume model.Volume) error {
	return nil
}
