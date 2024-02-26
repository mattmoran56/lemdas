import { getOptions } from "../../apiHelper";

const GetStaredDataset = async (datasetId) => {
	const requestOptions = getOptions("GET");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/dataset/${datasetId}/stared`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to get dataset stared status");
	}

	return response.json();
};

export default GetStaredDataset;
