import { getOptions } from "../../apiHelper";

const getAccess = async (datasetId) => {
	const requestOptions = getOptions("GET");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/dataset/${datasetId}/access`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to get dataset");
	}

	return response.json();
};

export default getAccess;
