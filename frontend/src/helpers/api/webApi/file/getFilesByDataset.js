import { getOptions } from "../../apiHelper";

const getDatasetFiles = async (datasetId) => {
	const requestOptions = getOptions("GET");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/dataset/${datasetId}/files`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to get files");
	}

	const data = await response.json();
	return data.files;
};

export default getDatasetFiles;
