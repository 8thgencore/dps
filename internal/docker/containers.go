package docker

import (
	"context"
	"dps/internal/table"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

const (
	containerID string = "container_id"
	nameID      string = "name_id"
	statusID    string = "status_id"
	imageID     string = "image_id"
	portID      string = "port_id"
	repoTagsID  string = "repo_tags_id"
	tagID       string = "tag_id"
	sizeID      string = "size_id"
	networkID   string = "network_id"
	createdID   string = "created_id"
	volumeID    string = "volume_id"
	driverID    string = "string_id"
)

func ShowContainers(ctx context.Context, cli *client.Client) {
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// Определите структуры Column для каждой колонки
	columns := []table.Column{
		{ID: containerID, Name: "Container ID", SortIndex: 0, Width: 12},
		{ID: nameID, Name: "Name", SortIndex: 1, Width: 40},
		{ID: statusID, Name: "Status", SortIndex: 2, Width: 11},
		{ID: imageID, Name: "Image", SortIndex: 3, Width: 40},
		{ID: portID, Name: "Ports", SortIndex: 4, Width: 20},
	}

	rowData := make([]table.RowData, len(containers))

	for i, container := range containers {
		data := make([]string, len(columns))
		for j, col := range columns {
			switch col.ID {
			case containerID:
				data[j] = container.ID[:12]
			case nameID:
				data[j] = container.Names[0][1:]
			case statusID:
				data[j] = container.Status
			case imageID:
				data[j] = container.Image
			case portID:
				portsList := make([]string, len(container.Ports))
				for _, port := range container.Ports {
					portsList = append(portsList, fmt.Sprintf("%d:%d", port.PublicPort, port.PrivatePort))
				}
				data[j] = sliteToString(portsList)
			}
		}
		rowData[i] = table.RowData{Data: data}
	}

	table.RenderTable(ctx, columns, rowData)

}

// ShowImages выводит информацию об изображениях Docker.
func ShowImages(ctx context.Context, cli *client.Client) {
	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	columns := []table.Column{
		{ID: imageID, Name: "Image ID", SortIndex: 0, Width: 12},
		{ID: repoTagsID, Name: "Repository Tags", SortIndex: 1, Width: 44},
		{ID: tagID, Name: "Tag", SortIndex: 3, Width: 15},
		{ID: createdID, Name: "Created", SortIndex: 2, Width: 20},
		{ID: sizeID, Name: "Size", SortIndex: 2, Width: 22},
		// Добавьте другие столбцы по вашему усмотрению.
	}

	rowData := make([]table.RowData, len(images))

	for i, image := range images {
		data := make([]string, len(columns))
		if len(image.RepoTags) > 0 {
			for j, col := range columns {
				switch col.ID {
				case imageID:
					data[j] = strings.Split(image.ID, ":")[1][:12]
				case repoTagsID:
					repoTags := []string{}
					for _, repoTag := range image.RepoTags {
						repo, _ := parseRepoTag(repoTag)
						repoTags = append(repoTags, repo)
					}
					data[j] = strings.Join(repoTags, ", ")
				case tagID:
					repoTags := []string{}
					for _, repoTag := range image.RepoTags {
						_, tag := parseRepoTag(repoTag)
						repoTags = append(repoTags, tag)
					}
					data[j] = strings.Join(repoTags, ", ")
				case createdID:
					createdTime := time.Unix(image.Created, 0)
					data[j] = createdTime.Format("2006-01-02 15:04:05")
				case sizeID:
					data[j] = formatBytes(float64(image.Size))
				}
			}
			rowData[i] = table.RowData{Data: data}
		}
	}

	table.RenderTable(ctx, columns, rowData)
}

// ShowNetworks выводит информацию о сетях Docker.
func ShowNetworks(ctx context.Context, cli *client.Client) {
	networks, err := cli.NetworkList(ctx, types.NetworkListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	columns := []table.Column{
		{ID: networkID, Name: "Network ID", SortIndex: 0, Width: 12},
		{ID: nameID, Name: "Name", SortIndex: 1, Width: 48},
		{ID: driverID, Name: "Driver", SortIndex: 2, Width: 15},
	}

	rowData := make([]table.RowData, len(networks))

	for i, network := range networks {
		data := make([]string, len(columns))
		for j, col := range columns {
			switch col.ID {
			case networkID:
				data[j] = network.ID[:12]
			case nameID:
				data[j] = network.Name
			case driverID:
				data[j] = network.Driver
				// Добавьте обработку других столбцов, если необходимо.
			}
		}
		rowData[i] = table.RowData{Data: data}
	}

	table.RenderTable(ctx, columns, rowData)
}

// ShowVolumes выводит информацию о томах Docker.
func ShowVolumes(ctx context.Context, cli *client.Client) {
	volumes, err := cli.VolumeList(ctx, volume.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	columns := []table.Column{
		{ID: driverID, Name: "Driver", SortIndex: 2, Width: 15},
		{ID: nameID, Name: "Name", SortIndex: 1, Width: 65},
	}

	rowData := make([]table.RowData, len(volumes.Volumes))

	for i, volume := range volumes.Volumes {
		data := make([]string, len(columns))
		for j, col := range columns {
			switch col.ID {
			case driverID:
				data[j] = volume.Driver
			case nameID:
				data[j] = volume.Name
			}
		}
		rowData[i] = table.RowData{Data: data}
	}

	table.RenderTable(ctx, columns, rowData)
}
