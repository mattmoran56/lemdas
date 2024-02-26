import { getOptions } from "../../apiHelper";

const updateStaredDataset = async (datasetId) => {
	const requestOptions = getOptions("POST");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/dataset/${datasetId}/stared`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to update dataset");
	}

	return response.json();
};

export default updateStaredDataset;
