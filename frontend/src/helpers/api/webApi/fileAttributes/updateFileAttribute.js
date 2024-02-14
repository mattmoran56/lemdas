import { getOptionsWithBody } from "../../apiHelper";

const updateFileAttribute = async (fileId, attributeId, attributeName, attributeValue) => {
	const data = {
		attribute_id: attributeId,
		attribute_name: attributeName,
		attribute_value: attributeValue,
	};

	const requestOptions = getOptionsWithBody("PUT", JSON.stringify(data));

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/file/${fileId}/attribute/`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to update attributes");
	}

	return response.json();
};

export default updateFileAttribute;
