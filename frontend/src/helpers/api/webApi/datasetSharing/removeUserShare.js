import { getOptions } from "../../apiHelper";

const removeUserShare = async (datasetId, userId) => {
	const requestOptions = getOptions("DELETE");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/dataset/${datasetId}/share/user/${userId}`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to delete user share");
	}

	return true;
};

export default removeUserShare;
