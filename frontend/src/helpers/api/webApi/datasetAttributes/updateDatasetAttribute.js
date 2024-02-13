import { getOptionsWithBody } from "../../apiHelper";

const updateDatasetAttribute = async (datasetId, attributeId, attributeName, attributeValue) => {
	const data = {
		attribute_id: attributeId,
		attribute_name: attributeName,
		attribute_value: attributeValue,
	};

	const requestOptions = getOptionsWithBody("PUT", JSON.stringify(data));

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/dataset/${datasetId}/attribute/`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to update attributes");
	}

	return response.json();
};

export default updateDatasetAttribute;
