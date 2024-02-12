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

	return response.json();
};

export default getDatasetFiles;
