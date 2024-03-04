import { getOptions } from "../../apiHelper";

const deleteGroupMember = async (groupId, userId) => {
	const requestOptions = getOptions("DELETE");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/group/${groupId}/member/${userId}`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to remove member from group");
	}

	return true;
};

export default deleteGroupMember;
