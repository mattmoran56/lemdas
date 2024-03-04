import React, { useEffect, useState } from "react";
import { PlusIcon } from "@heroicons/react/24/outline";

import getUserShares from "../../../helpers/api/webApi/datasetSharing/getUserShares";
import getGroupShares from "../../../helpers/api/webApi/datasetSharing/getGroupShares";
import ErrorToast from "../../../helpers/toast/errorToast";
import User from "../../basic/User";
import Group from "../../basic/Group";
import UserShareModal from "./UserShareModal";
import removeUserShare from "../../../helpers/api/webApi/datasetSharing/removeUserShare";
import removeGroupShare from "../../../helpers/api/webApi/datasetSharing/removeGroupShare";
import GroupShareModal from "./GroupShareModal";

const Sharing = ({ datasetId, writeAccess }) => {
	const [users, setUsers] = useState([]);
	const [groups, setGroups] = useState([]);

	const [showUserModal, setShowUserModal] = useState(false);
	const [showGroupModal, setShowGroupModal] = useState(false);

	const handleRemoveUserShare = (userId) => {
		removeUserShare(datasetId, userId).then(() => {
			getUserShares(datasetId).then((data) => {
				setUsers(data.users);
			}).catch((error) => {
				ErrorToast(error);
			});
		}).catch((error) => {
			ErrorToast(error);
		});
	};

	const handleRemoveGroupShare = (groupId) => {
		removeGroupShare(datasetId, groupId).then(() => {
			getGroupShares(datasetId).then((data) => {
				setGroups(data.groups);
			}).catch((error) => {
				ErrorToast(error);
			});
		}).catch((error) => {
			ErrorToast(error);
		});
	};

	useEffect(() => {
		getUserShares(datasetId).then((data) => {
			setUsers(data.users);
		}).catch((error) => {
			ErrorToast(error);
		});
		getGroupShares(datasetId).then((data) => {
			setGroups(data.groups);
		}).catch((error) => {
			ErrorToast(error);
		});
	}, []);

	return (
		<div className="w-full">
			{writeAccess
				? (
					<>
						<UserShareModal
							isOpen={showUserModal}
							setIsOpen={setShowUserModal}
							datasetId={datasetId}
							onClose={() => {
								getUserShares(datasetId).then((data) => {
									setUsers(data.users);
								}).catch((error) => {
									ErrorToast(error);
								});
							}}
						/>
						<GroupShareModal
							isOpen={showGroupModal}
							setIsOpen={setShowGroupModal}
							datasetId={datasetId}
							onClose={() => {
								getGroupShares(datasetId).then((data) => {
									setGroups(data.groups);
								}).catch((error) => {
									ErrorToast(error);
								});
							}}
						/>
					</>
				) : null}

			{!writeAccess && users.length === 0
				? null
				: (
					<div className="mb-4">
						<h3 className="text-lg font-semibold">Users shared with</h3>
						<div className="flex flex-wrap w-full">
							{users.map((user) => {
								return (
									<User
										key={user.id}
										name={`${user.user.first_name} ${user.user.last_name}`}
										avatar={user.avatar}
										access={user.write_access ? "write" : "read"}
										onRemove={() => {
											handleRemoveUserShare(user.user.id);
										}}
										writeAccess={writeAccess}
									/>
								);
							})}
							{writeAccess
								? (
									<button
										type="button"
										aria-label="add user"
										className={`w-44 h-36 border-2 border-oxfordblue-200 border-dashed my-2 mr-4 
													p-4 rounded-md flex flex-col items-center hover:!underline`}
										onClick={() => {
											setShowUserModal(true);
										}}
									>
										<div
											className={`rounded-full border-2 border-dashed border-gray-400 
														aspect-square p-4`}
										>
											<PlusIcon className="h-6 w-6 text-oxfordblue-400" />
										</div>
										<p className="text-md font-semibold py-2">Share with user</p>
									</button>
								) : null}
						</div>
					</div>
				)}
			{!writeAccess && groups.length === 0
				? null
				: (
					<div>
						<h3 className="text-lg font-semibold mt-4">Groups shared with</h3>
						<div className="flex flex-wrap w-full">
							{groups.map((group) => {
								return (
									<Group
										key={group.id}
										name={group.group.group_name}
										access={group.write_access ? "write" : "read"}
										onRemove={() => {
											handleRemoveGroupShare(group.group.id);
										}}
										writeAccess={writeAccess}
									/>
								);
							})}
							{writeAccess
								? (
									<button
										type="button"
										aria-label="add user"
										className={`w-44 h-36 border-2 border-oxfordblue-200 border-dashed my-2 mr-4
													p-4 rounded-md flex flex-col items-center hover:!underline`}
										onClick={() => {
											setShowGroupModal(true);
										}}
									>
										<div
											className={`rounded-full border-2 border-dashed border-gray-400
											 			aspect-square p-4`}
										>
											<PlusIcon className="h-6 w-6 text-oxfordblue-400" />
										</div>
										<p className="text-md font-semibold py-2">Share with Group</p>
									</button>
								) : null}
						</div>
					</div>
				)}
		</div>
	);
};

export default Sharing;
