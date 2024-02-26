import { getOptions } from "../../apiHelper";

const deleteDatasetAttribute = async (datasetId, attributeId) => {
	const requestOptions = getOptions("DELETE");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/dataset/${datasetId}/attribute/${attributeId}`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to delete attributes");
	}

	return true;
};

export default deleteDatasetAttribute;
