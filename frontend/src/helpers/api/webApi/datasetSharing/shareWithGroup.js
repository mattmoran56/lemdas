import { getOptionsWithBody } from "../../apiHelper";

const shareWithGroup = async (datasetId, groupId, writeAccess) => {
	const requestOptions = getOptionsWithBody("POST", JSON.stringify({ group_id: groupId, write_access: writeAccess }));

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/dataset/${datasetId}/share/group`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to share with group");
	}

	return response.json();
};

export default shareWithGroup;
