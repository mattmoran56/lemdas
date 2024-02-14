import { getOptionsWithBody } from "../../apiHelper";

const addFileAttribute = async (fileId, attributeName, attributeValue) => {
	const body = JSON.stringify({
		attribute_name: attributeName,
		attribute_value: attributeValue,
	});
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
