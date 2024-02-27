import { getOptionsWithBody } from "../../apiHelper";

const addCollaborator = async (datasetId, userId) => {
	const body = JSON.stringify({
		user_id: userId,
	});
	const requestOptions = getOptionsWithBody("POST", body);
	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/dataset/${datasetId}/collaborator`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to add collaborator");
	}

	return response.json();
};

export default addCollaborator;
