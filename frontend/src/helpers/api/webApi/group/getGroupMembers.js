import { getOptions } from "../../apiHelper";

const getGroupMembers = async (groupId) => {
	const requestOptions = getOptions("GET");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/group/${groupId}/member`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to get group members");
	}

	return response.json();
};

export default getGroupMembers;
