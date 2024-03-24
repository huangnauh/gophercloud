/*
Package servers provides information and interaction with the server API
resource in the OpenStack Compute service.

A server is a virtual machine instance in the compute system. In order for
one to be provisioned, a valid flavor and image are required.

Example to List Servers

	listOpts := servers.ListOpts{
		AllTenants: true,
	}

	allPages, err := servers.ListSimple(computeClient, listOpts).AllPages(context.TODO())
	if err != nil {
		panic(err)
	}

	allServers, err := servers.ExtractServers(allPages)
	if err != nil {
		panic(err)
	}

	for _, server := range allServers {
		fmt.Printf("%+v\n", server)
	}

Example to List Detail Servers

	listOpts := servers.ListOpts{
		AllTenants: true,
	}

	allPages, err := servers.List(computeClient, listOpts).AllPages(context.TODO())
	if err != nil {
		panic(err)
	}

	allServers, err := servers.ExtractServers(allPages)
	if err != nil {
		panic(err)
	}

	for _, server := range allServers {
		fmt.Printf("%+v\n", server)
	}

Example to Create a Server

	createOpts := servers.CreateOpts{
		Name:      "server_name",
		ImageRef:  "image-uuid",
		FlavorRef: "flavor-uuid",
	}

	server, err := servers.Create(context.TODO(), computeClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a Server

	serverID := "d9072956-1560-487c-97f2-18bdf65ec749"
	err := servers.Delete(context.TODO(), computeClient, serverID).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to Force Delete a Server

	serverID := "d9072956-1560-487c-97f2-18bdf65ec749"
	err := servers.ForceDelete(context.TODO(), computeClient, serverID).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to Reboot a Server

	rebootOpts := servers.RebootOpts{
		Type: servers.SoftReboot,
	}

	serverID := "d9072956-1560-487c-97f2-18bdf65ec749"

	err := servers.Reboot(context.TODO(), computeClient, serverID, rebootOpts).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to Rebuild a Server

	rebuildOpts := servers.RebuildOpts{
		Name:    "new_name",
		ImageID: "image-uuid",
	}

	serverID := "d9072956-1560-487c-97f2-18bdf65ec749"

	server, err := servers.Rebuilt(computeClient, serverID, rebuildOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Resize a Server

	resizeOpts := servers.ResizeOpts{
		FlavorRef: "flavor-uuid",
	}

	serverID := "d9072956-1560-487c-97f2-18bdf65ec749"

	err := servers.Resize(context.TODO(), computeClient, serverID, resizeOpts).ExtractErr()
	if err != nil {
		panic(err)
	}

	err = servers.ConfirmResize(context.TODO(), computeClient, serverID).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to Snapshot a Server

	snapshotOpts := servers.CreateImageOpts{
		Name: "snapshot_name",
	}

	serverID := "d9072956-1560-487c-97f2-18bdf65ec749"

	image, err := servers.CreateImage(context.TODO(), computeClient, serverID, snapshotOpts).ExtractImageID()
	if err != nil {
		panic(err)
	}
*/
package servers
