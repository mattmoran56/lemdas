import React, { useEffect, useState } from "react";
import { PlusIcon } from "@heroicons/react/24/outline";

import getUsersGroups from "../../../helpers/api/webApi/group/getUsersGroups";
import ErrorToast from "../../../helpers/toast/errorToast";
import Group from "../../basic/Group";
import GroupInfoModal from "./GroupInfoModal";
import NewGroupModal from "./NewGroupModal";

const GroupButton = ({ group, refreshGroups }) => {
	const [isOpen, setIsOpen] = useState(false);

	return (
		<>
			<button
				type="button"
				aria-label={group.group_name}
				onClick={() => { setIsOpen(true); }}
			>
				<Group
					name={group.group_name}
					access={group.write_access ? "write" : "read"}
					onRemove={() => {}}
					noAccess
				/>
			</button>
			<GroupInfoModal
				isOpen={isOpen}
				setIsOpen={setIsOpen}
				group={group}
				onClose={refreshGroups}
			/>
		</>
	);
};

const Groups = () => {
	const [groups, setGroups] = useState([]);

	const [showNewGroupModal, setShowNewGroupModal] = useState(false);

	useEffect(() => {
		getUsersGroups().then((data) => {
			setGroups(data.groups);
		}).catch((error) => {
			ErrorToast(error);
		});
	}, []);

	return (
		<div className="w-full pt-8">
			<NewGroupModal
				isOpen={showNewGroupModal}
				setIsOpen={setShowNewGroupModal}
				onClose={() => {
					getUsersGroups().then((data) => {
						setGroups(data.groups);
					}).catch((error) => {
						ErrorToast(error);
					});
				}}
			/>
			<h1 className="text-2xl font-semibold">Groups</h1>
			<div className="w-full h-[2px] bg-oxfordblue" />
			<div className="w-full flex flex-wrap">
				{groups.map((group) => {
					return (
						<GroupButton
							group={group}
							key={group.id}
							refreshGroups={() => {
								getUsersGroups().then((data) => {
									setGroups(data.groups);
								}).catch((error) => {
									ErrorToast(error);
								});
							}}
						/>
					);
				})}
				<button
					type="button"
					aria-label="create group"
					className={`w-44 h-36 border-2 border-oxfordblue-200 border-dashed my-2 mr-4 p-4 rounded-md flex
								flex-col items-center`}
					onClick={() => { setShowNewGroupModal(true); }}
				>
					<div
						className={`border-dashed border-2 border-gray-400 rounded-full aspect-square p-4
									hover:!underline`}
					>
						<PlusIcon className="h-6 w-6 text-oxfordblue-400" />
					</div>
					<p className="text-md font-semibold py-2">New group</p>
				</button>
			</div>
		</div>
	);
};

export default Groups;
