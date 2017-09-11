package endpoint

import (
	"github.com/kassisol/hbm/allow"
	"github.com/kassisol/hbm/docker/dcb"
	"github.com/kassisol/hbm/pkg/uri"
)

func GetUris() *uri.URIs {
	uris := uri.New()

	uris.Register("GET", `^/containers/json`, allow.AllowTrue, dcb.ContainerList, "container_list", "container ls", "List containers")
	uris.Register("POST", `^/containers/create`, allow.AllowContainerCreate, dcb.ContainerCreate, "container_create", "container create", "Create a new container")
	uris.Register("GET", `^/containers/(.+)/json`, allow.AllowTrue, dcb.ContainerInspect, "container_inspect", "container inspect", "Display detailed information on one or more containers")
	uris.Register("GET", `^/containers/(.+)/top`, allow.AllowTrue, dcb.ContainerTop, "container_top", "container top", "Display the running processes of a container")
	uris.Register("GET", `^/containers/(.+)/logs`, allow.AllowTrue, dcb.ContainerLogs, "container_logs", "container logs", "Fetch the logs of a container")
	uris.Register("GET", `^/containers/(.+)/changes`, allow.AllowTrue, dcb.ContainerChanges, "container_changes", "events", "Get real time events from the server")
	uris.Register("GET", `^/containers/(.+)/export`, allow.AllowTrue, dcb.ContainerExport, "container_export", "container export", "Export a container's filesystem as a tar archive")
	uris.Register("GET", `^/containers/(.+)/stats`, allow.AllowTrue, dcb.ContainerStats, "container_stats", "container stats", "Display a live stream of container(s) resource usage statistics")
	uris.Register("POST", `^/containers/(.+)/resize`, allow.AllowTrue, dcb.ContainerResize, "container_resize", "resize", "Resize a container TTY")
	uris.Register("POST", `^/containers/(.+)/start`, allow.AllowTrue, dcb.ContainerStart, "container_start", "container start", "Start one or more stopped containers")
	uris.Register("POST", `^/containers/(.+)/stop`, allow.AllowTrue, dcb.ContainerStop, "container_stop", "container stop", "Stop one or more running container")
	uris.Register("POST", `^/containers/(.+)/restart`, allow.AllowTrue, dcb.ContainerRestart, "container_restart", "container restart", "Restart one or more containers")
	uris.Register("POST", `^/containers/(.+)/kill`, allow.AllowTrue, dcb.ContainerKill, "container_kill", "container kill", "Kill one or more running containers")
	uris.Register("POST", `^/containers/(.+)/update`, allow.AllowTrue, dcb.ContainerUpdate, "container_update", "container update", "Update configuration of one or more containers")
	uris.Register("POST", `^/containers/(.+)/rename`, allow.AllowTrue, dcb.ContainerRename, "container_rename", "container rename", "Rename a container")
	uris.Register("POST", `^/containers/(.+)/pause`, allow.AllowTrue, dcb.ContainerPause, "container_pause", "container pause", "Pause all processes within one or more containers")
	uris.Register("POST", `^/containers/(.+)/unpause`, allow.AllowTrue, dcb.ContainerUnpause, "container_unpause", "container unpause", "Unpause all processes within one or more containers")
	uris.Register("POST", `^/containers/(.+)/attach`, allow.AllowTrue, dcb.ContainerAttach, "container_attach", "container attach", "Attach to a running container")
	uris.Register("GET", `^/containers/(.+)/attach/ws`, allow.AllowTrue, dcb.ContainerAttachWS, "container_attach_ws", "attach_ws", "Attach to a running container (websocket)")
	uris.Register("POST", `^/containers/(.+)/wait`, allow.AllowTrue, dcb.ContainerWait, "container_wait", "container wait", "Block until one or more containers stop, then print their exit codes")
	uris.Register("DELETE", `^/containers/(.+)`, allow.AllowTrue, dcb.ContainerRemove, "container_remove", "container rm", "Remove one or more containers")
	uris.Register("HEAD", `^/containers/(.+)/archive`, allow.AllowTrue, dcb.ContainerArchiveInfo, "container_archive_info", "archive", "Retrieving information about files and folders in a container")
	uris.Register("GET", `^/containers/(.+)/archive`, allow.AllowTrue, dcb.ContainerArchive, "container_archive", "archive", "Get an archive of a filesystem resource in a container")
	uris.Register("PUT", `^/containers/(.+)/archive`, allow.AllowTrue, dcb.ContainerArchiveExtract, "container_archive_extract", "archive", "Extract an archive of files or folders to a directory in a container")
	uris.Register("POST", `^/containers/prune`, allow.AllowTrue, dcb.ContainerPrune, "container_prune", "container prune", "Remove all stopped containers")

	uris.Register("GET", `^/images/json`, allow.AllowTrue, dcb.ImageList, "image_list", "image ls", "List images")
	uris.Register("POST", `^/build`, allow.AllowTrue, dcb.ImageBuild, "image_build", "image build", "Build an image from a Dockerfile")
	uris.Register("POST", `^/images/create`, allow.AllowImageCreate, dcb.ImageCreate, "image_create", "image pull", "Pull an image or a repository from a registry")
	uris.Register("GET", `^/images/(.+)/json`, allow.AllowTrue, dcb.ImageInspect, "image_inspect", "image inspect", "Display detailed information on one or more images")
	uris.Register("GET", `^/images/(.+)/history`, allow.AllowTrue, dcb.ImageHistory, "image_history", "image history", "Show the history of an image")
	uris.Register("POST", `^/images/(.+)/push`, allow.AllowTrue, dcb.ImagePush, "image_push", "image push", "Push an image or a repository to a registry")
	uris.Register("POST", `^/images/(.+)/tag`, allow.AllowTrue, dcb.ImageTag, "image_tag", "image tag", "Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE")
	uris.Register("DELETE", `^/images/(.+)`, allow.AllowTrue, dcb.ImageRemove, "image_remove", "image rm", "Remove one or more images")
	uris.Register("GET", `^/images/search`, allow.AllowTrue, dcb.ImageSearch, "image_search", "search", "Search the Docker Hub for images")
	uris.Register("POST", `^/images/prune`, allow.AllowTrue, dcb.ImagePrune, "image_prune", "image prune", "Delete unused images")
	uris.Register("POST", `^/commit`, allow.AllowTrue, dcb.Commit, "commit", "commit", "Create a new image from a container's changes")
	uris.Register("GET", `^/images/(.+)/get`, allow.AllowTrue, dcb.ImageSaveImage, "image_save_image", "image save", "Save one or more images to a tar archive")
	uris.Register("GET", `^/images/get`, allow.AllowTrue, dcb.ImageSaveImages, "image_save_images", "image save", "Save one or more images to a tar archive")
	uris.Register("POST", `^/images/load`, allow.AllowTrue, dcb.ImageLoad, "image_load", "image load", "Load an image from a tar archive or STDIN")

	uris.Register("GET", `^/networks$`, allow.AllowTrue, dcb.NetworkList, "network_list", "network ls", "List all networks")
	uris.Register("GET", `^/networks/(.+)`, allow.AllowTrue, dcb.NetworkInspect, "network_inspect", "network inspect", "Display detailed network information")
	uris.Register("DELETE", `^/networks/(.+)`, allow.AllowTrue, dcb.NetworkRemove, "network_remove", "network rm", "Remove a network")
	uris.Register("POST", `^/networks/create`, allow.AllowTrue, dcb.NetworkCreate, "network_create", "network create", "Create a network")
	uris.Register("POST", `^/networks/(.+)/connect`, allow.AllowTrue, dcb.NetworkConnect, "network_connect", "network connect", "Connect container to a network")
	uris.Register("POST", `^/networks/(.+)/disconnect`, allow.AllowTrue, dcb.NetworkDisconnect, "network_disconnect", "network disconnect", "Disconnect container from a network")
	uris.Register("POST", `^/networks/prune`, allow.AllowTrue, dcb.NetworkPrune, "network_prune", "network prune", "Delete unused networks")

	uris.Register("GET", `^/volumes$`, allow.AllowTrue, dcb.VolumeList, "volume_list", "volume ls", "List volumes")
	uris.Register("POST", `^/volumes/create`, allow.AllowTrue, dcb.VolumeCreate, "volume_create", "volume create", "Create a volume")
	uris.Register("GET", `^/volumes/(.+)`, allow.AllowTrue, dcb.VolumeInspect, "volume_inspect", "volume inspect", "Return low-level information on a volume")
	uris.Register("DELETE", `^/volumes/(.+)`, allow.AllowTrue, dcb.VolumeRemove, "volume_remove", "volume rm", "Remove a volume")
	uris.Register("POST", `^/volumes/prune`, allow.AllowTrue, dcb.VolumePrune, "volume_prune", "volume prune", "Delete unused volumes")

	uris.Register("POST", `^/containers/(.+)/exec`, allow.AllowTrue, dcb.ContainerExecCreate, "container_exec_create", "exec", "Run a command in a running container")
	uris.Register("POST", `^/exec/(.+)/start`, allow.AllowTrue, dcb.ExecStart, "exec_start", "exec", "Exec Start")
	uris.Register("POST", `^/exec/(.+)/resize`, allow.AllowTrue, dcb.ExecResize, "exec_resize", "exec", "Exec Resize")
	uris.Register("GET", `^/exec/(.+)/json`, allow.AllowTrue, dcb.ExecInspect, "exec_inspect", "exec", "Exec Inspect")

	uris.Register("GET", `^/swarm`, allow.AllowTrue, dcb.SwarmInspect, "swarm_inspect", "swarm info", "Get swarm info")
	uris.Register("POST", `^/swarm/init`, allow.AllowTrue, dcb.SwarmInit, "swarm_init", "swarm init", "Initialize a new swarm")
	uris.Register("POST", `^/swarm/join`, allow.AllowTrue, dcb.SwarmJoin, "swarm_join", "swarm join", "Join an existing swarm")
	uris.Register("POST", `^/swarm/leave`, allow.AllowTrue, dcb.SwarmLeave, "swarm_leave", "swarm leave", "Leave a swarm")
	uris.Register("POST", `^/swarm/update`, allow.AllowTrue, dcb.SwarmUpdate, "swarm_update", "swarm update", "Update a swarm")
	uris.Register("GET", `^/swarm/unlockkey`, allow.AllowTrue, dcb.SwarmUnlockKey, "swarm_unlock_key", "swarm unlock", "Get the unlock key")
	uris.Register("POST", `^/swarm/unlock`, allow.AllowTrue, dcb.SwarmUnlock, "swarm_unlock", "swarm unlock", "Unlock a locked manager")

	uris.Register("GET", `^/nodes`, allow.AllowTrue, dcb.NodeList, "node_list", "node ls", "List nodes")
	uris.Register("GET", `^/nodes/(.+)`, allow.AllowTrue, dcb.NodeInspect, "node_inspect", "node inspect", "Return low-level information on the node id")
	uris.Register("DELETE", `^/nodes/(.+)`, allow.AllowTrue, dcb.NodeRemove, "node_remove", "node rm", "Remove a node [id] from the swarm")
	uris.Register("POST", `^/nodes/(.+)/update`, allow.AllowTrue, dcb.NodeUpdate, "node_update", "node update", "Update the node id")

	uris.Register("GET", `^/services`, allow.AllowTrue, dcb.ServiceList, "service_list", "service ls", "List services")
	uris.Register("POST", `^/services/create`, allow.AllowServiceCreate, dcb.ServiceCreate, "service_create", "service create", "Create a service")
	uris.Register("GET", `^/services/(.+)`, allow.AllowTrue, dcb.ServiceInspect, "service_inspect", "service inspect", "Inspect a service")
	uris.Register("DELETE", `^/services/(.+)`, allow.AllowTrue, dcb.ServiceRemove, "service_remove", "service rm", "Remove a service")
	uris.Register("POST", `^/services/(.+)/update`, allow.AllowServiceCreate, dcb.ServiceUpdate, "service_update", "service update", "Update a service")
	uris.Register("GET", `^/services/(.+)/logs`, allow.AllowTrue, dcb.ServiceLogs, "service_logs", "service logs", "Get service logs")

	uris.Register("GET", `^/tasks`, allow.AllowTrue, dcb.TaskList, "task_list", "stask services", "List tasks")
	uris.Register("GET", `^/tasks/(.+)`, allow.AllowTrue, dcb.TaskInspect, "task_inspect", "stask tasks", "Get details on a task")

	uris.Register("GET", `^/secrets`, allow.AllowTrue, dcb.SecretList, "secret_list", "secret ls", "List secrets")
	uris.Register("POST", `^/secrets/create`, allow.AllowTrue, dcb.SecretCreate, "secret_create", "secret create", "Create a secret")
	uris.Register("GET", `^/secrets/(.+)`, allow.AllowTrue, dcb.SecretInspect, "secret_inspect", "secret inspect", "Inspect a secret")
	uris.Register("DELETE", `^/secrets/(.+)`, allow.AllowTrue, dcb.SecretRemove, "secret_remove", "secret rm", "Delete a secret")
	uris.Register("POST", `^/secrets/(.+)/update`, allow.AllowTrue, dcb.SecretUpdate, "secret_update", "secret update", "Update a secret")

	uris.Register("GET", `^/plugins`, allow.AllowTrue, dcb.PluginList, "plugin_list", "plugin ls", "List plugins")
	uris.Register("GET", `^/plugins/privileges`, allow.AllowTrue, dcb.PluginPrivileges, "plugin_privileges", "plugin ls", "Get plugin privileges")
	uris.Register("POST", `^/plugins/pull`, allow.AllowTrue, dcb.PluginPull, "plugin_pull", "plugin pull", "Install a plugin")
	uris.Register("GET", `^/plugins/(.+)/json`, allow.AllowTrue, dcb.PluginInspect, "plugin_inspect", "plugin inspect", "Inspect a plugin")
	uris.Register("DELETE", `^/plugins/(.+)`, allow.AllowTrue, dcb.PluginRemove, "plugin_remove", "plugin rm", "Remove a plugin")
	uris.Register("POST", `^/plugins/(.+)/enable`, allow.AllowTrue, dcb.PluginEnable, "plugin_enable", "plugin enable", "Enable a plugin")
	uris.Register("POST", `^/plugins/(.+)/disable`, allow.AllowTrue, dcb.PluginDisable, "plugin_disable", "plugin disable", "Disable a plugin")
	uris.Register("POST", `^/plugins/(.+)/upgrade`, allow.AllowTrue, dcb.PluginUpgrade, "plugin_upgrade", "plugin upgrade", "Upgrade a plugin")
	uris.Register("POST", `^/plugins/create`, allow.AllowTrue, dcb.PluginCreate, "plugin_create", "plugin create", "Create a plugin")
	uris.Register("POST", `^/plugins/(.+)/push`, allow.AllowTrue, dcb.PluginPush, "plugin_push", "plugin push", "Push a plugin")
	uris.Register("POST", `^/plugins/(.+)/set`, allow.AllowTrue, dcb.PluginSet, "plugin_set", "plugin set", "Configure a plugin")

	uris.Register("POST", `^/auth`, allow.AllowTrue, dcb.Auth, "auth", "login", "Log in to a Docker registry")
	uris.Register("GET", `^/info`, allow.AllowTrue, dcb.Info, "info", "info", "Display system-wide information")
	uris.Register("GET", `^/version`, allow.AllowTrue, dcb.Version, "version", "version", "Show the Docker version information")
	uris.Register("GET", `^/_ping`, allow.AllowTrue, dcb.Ping, "ping", "*ping*", "Ping the docker server")
	uris.Register("GET", `^/events`, allow.AllowTrue, dcb.Events, "events", "events", "Monitor Docker’s events")
	uris.Register("GET", `^/system/df`, allow.AllowTrue, dcb.SystemDF, "system_df", "system df", "Get data usage information")

	uris.Register("GET", `^/configs`, allow.AllowTrue, dcb.ConfigList, "config_list", "config ls", "List configs")
	uris.Register("POST", `^/configs/create`, allow.AllowTrue, dcb.ConfigCreate, "config_create", "config create", "Create a config")
	uris.Register("GET", `^/configs/(.+)`, allow.AllowTrue, dcb.ConfigInspect, "config_inspect", "config inspect", "Inspect a config")
	uris.Register("DELETE", `^/configs/(.+)`, allow.AllowTrue, dcb.ConfigRemove, "config_remove", "config rm", "Delete a config")
	uris.Register("POST", `^/configs/(.+)/update`, allow.AllowTrue, dcb.ConfigUpdate, "config_update", "config update", "Update a config")

	uris.Register("GET", `^/distribution/(.+)/json`, allow.AllowTrue, dcb.DistributionInfo, "distribution_info", "", "Get image information from the registry")

	uris.Register("GET", `^/tasks/(.+)/logs`, allow.AllowTrue, dcb.TaskLogs, "task_logs", "task logs", "Get task logs")

	uris.Register("OPTIONS", `^/(.*)`, allow.AllowTrue, dcb.Anyroute, "anyroute_options", "", "Anyroute OPTIONS")

	return uris
}
