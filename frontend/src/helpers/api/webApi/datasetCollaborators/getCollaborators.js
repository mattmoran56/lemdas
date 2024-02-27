import { getOptions } from "../../apiHelper";

const getCollaborators = async (datasetId) => {
	const requestOptions = getOptions("GET");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/dataset/${datasetId}/collaborator`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to get dataset collaborators");
	}

	return response.json();
};

export default getCollaborators;
