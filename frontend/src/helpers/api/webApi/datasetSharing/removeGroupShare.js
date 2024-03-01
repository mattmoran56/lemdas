import { getOptions } from "../../apiHelper";

const removeGroupShare = async (datasetId, groupId) => {
	const requestOptions = getOptions("DELETE");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/dataset/${datasetId}/share/group/${groupId}`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to delete group share");
	}

	return true;
};

export default removeGroupShare;
