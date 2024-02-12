import { getOptions } from "../../apiHelper";

const GetDataset = async (datasetId) => {
	const requestOptions = getOptions("GET");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/dataset/${datasetId}`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to get datasets");
	}

	return response.json();
};

export default GetDataset;
