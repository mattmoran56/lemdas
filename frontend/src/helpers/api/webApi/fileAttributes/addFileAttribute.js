import { getOptionsWithBody } from "../../apiHelper";

const addFileAttribute = async (fileId, attributeName, attributeValue, attributeGroupId) => {
	console.log(attributeGroupId);
	const body = JSON.stringify({
		attribute_name: attributeName,
		attribute_group_id: attributeGroupId,
		attribute_value: attributeValue,
	});
	console.log(body);
	const requestOptions = getOptionsWithBody("POST", body);
	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/file/${fileId}/attribute`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to add attributes");
	}

	return response.json();
};

export default addFileAttribute;
