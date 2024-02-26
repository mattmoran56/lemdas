import { getOptionsWithBody } from "../../apiHelper";

const updateDataset = async (datasetId, datasetName, isPublic, ownerId) => {
	const data = {
		dataset_name: datasetName,
		is_public: isPublic,
		owner_id: ownerId,
	};

	const requestOptions = getOptionsWithBody("PUT", JSON.stringify(data));

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/dataset/${datasetId}`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to update dataset");
	}

	return response.json();
};

export default updateDataset;
