/*
Package volumeactions provides information and interaction with volumes in the
OpenStack Block Storage service. A volume is a detachable block storage
device, akin to a USB hard drive.

Example of Attaching a Volume to an Instance

	attachOpts := volumeactions.AttachOpts{
		MountPoint:   "/mnt",
		Mode:         "rw",
		InstanceUUID: server.ID,
	}

	err := volumeactions.Attach(context.TODO(), client, volume.ID, attachOpts).ExtractErr()
	if err != nil {
		panic(err)
	}

	detachOpts := volumeactions.DetachOpts{
		AttachmentID: volume.Attachments[0].AttachmentID,
	}

	err = volumeactions.Detach(context.TODO(), client, volume.ID, detachOpts).ExtractErr()
	if err != nil {
		panic(err)
	}

Example of Creating an Image from a Volume

	uploadImageOpts := volumeactions.UploadImageOpts{
		ImageName: "my_vol",
		Force:     true,
	}

	volumeImage, err := volumeactions.UploadImage(context.TODO(), client, volume.ID, uploadImageOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", volumeImage)

Example of Extending a Volume's Size

	extendOpts := volumeactions.ExtendSizeOpts{
		NewSize: 100,
	}

	err := volumeactions.ExtendSize(context.TODO(), client, volume.ID, extendOpts).ExtractErr()
	if err != nil {
		panic(err)
	}

Example of Initializing a Volume Connection

	connectOpts := &volumeactions.InitializeConnectionOpts{
		IP:        "127.0.0.1",
		Host:      "stack",
		Initiator: "iqn.1994-05.com.redhat:17cf566367d2",
		Multipath: gophercloud.Disabled,
		Platform:  "x86_64",
		OSType:    "linux2",
	}

	connectionInfo, err := volumeactions.InitializeConnection(context.TODO(), client, volume.ID, connectOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", connectionInfo["data"])

	terminateOpts := &volumeactions.InitializeConnectionOpts{
		IP:        "127.0.0.1",
		Host:      "stack",
		Initiator: "iqn.1994-05.com.redhat:17cf566367d2",
		Multipath: gophercloud.Disabled,
		Platform:  "x86_64",
		OSType:    "linux2",
	}

	err = volumeactions.TerminateConnection(context.TODO(), client, volume.ID, terminateOpts).ExtractErr()
	if err != nil {
		panic(err)
	}

Example of Setting a Volume's Bootable status

	options := volumeactions.BootableOpts{
		Bootable: true,
	}

	err := volumeactions.SetBootable(context.TODO(), client, volume.ID, options).ExtractErr()
	if err != nil {
		panic(err)
	}

Example of Changing Type of a Volume

	changeTypeOpts := volumeactions.ChangeTypeOpts{
		NewType:         "ssd",
		MigrationPolicy: volumeactions.MigrationPolicyOnDemand,
	}

	err = volumeactions.ChangeType(context.TODO(), client, volumeID, changeTypeOpts).ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package volumeactions
