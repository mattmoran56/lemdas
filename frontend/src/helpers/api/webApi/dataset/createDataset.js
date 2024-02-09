import { getOptionsWithBody } from "../../apiHelper";

const createDataset = async (datasetName) => {
	const body = JSON.stringify({
		dataset_name: datasetName,
	});
	const requestOptions = getOptionsWithBody("POST", body);

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/dataset`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to create dataset");
	}

	return response.json();
};

export default createDataset;
