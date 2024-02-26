import { getOptions } from "../../apiHelper";

const deleteDataset = async (datasetId) => {
	const requestOptions = getOptions("DELETE");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/dataset/${datasetId}`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to delete dataset");
	}

	return true;
};

export default deleteDataset;
