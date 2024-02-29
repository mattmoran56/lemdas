import { getOptions } from "../../apiHelper";

const getUserShares = async (datasetId) => {
	const requestOptions = getOptions("GET");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/dataset/${datasetId}/share/user`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to get user shares");
	}

	return response.json();
};

export default getUserShares;
