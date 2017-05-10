package plugin

import "io/ioutil"

func (client *Client) DownloadPlugin(pluginURL string, path string) error {
	request, err := client.newGETRequest(pluginURL)
	if err != nil {
		return err
	}

	response := Response{}
	err = client.connection.Make(request, &response)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, response.RawResponse, 0700)
	if err != nil {
		return err
	}

	return nil
}
