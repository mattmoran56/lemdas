import { getOptionsWithBody } from "../../apiHelper";

const updateFile = async (fileId, name, ownerId, datasetId, status) => {
	const data = {
		name,
		owner_id: ownerId,
		dataset_id: datasetId,
		status,
	};

	const requestOptions = getOptionsWithBody("PUT", JSON.stringify(data));

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/file/${fileId}`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to update file");
	}

	return response.json();
};

export default updateFile;
