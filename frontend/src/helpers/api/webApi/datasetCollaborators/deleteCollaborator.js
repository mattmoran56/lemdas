import { getOptions } from "../../apiHelper";

const deleteCollaborator = async (datasetId, userId) => {
	const requestOptions = getOptions("DELETE");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/dataset/${datasetId}/collaborator/${userId}`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to delete collaborator");
	}

	return true;
};

export default deleteCollaborator;
