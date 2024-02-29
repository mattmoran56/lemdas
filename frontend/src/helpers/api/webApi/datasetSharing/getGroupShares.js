import { getOptions } from "../../apiHelper";

const getGroupShares = async (datasetId) => {
	const requestOptions = getOptions("GET");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/dataset/${datasetId}/share/group`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to get group shares");
	}

	return response.json();
};

export default getGroupShares;
