import React, { useEffect, useState } from "react";
import { ChevronDownIcon, PlusIcon, TrashIcon } from "@heroicons/react/24/outline";

import addFileAttribute from "../../helpers/api/webApi/fileAttributes/addFileAttribute";
import deleteFileAttribute from "../../helpers/api/webApi/fileAttributes/deleteFileAttribute";
import updateFileAttribute from "../../helpers/api/webApi/fileAttributes/updateFileAttribute";
import ErrorToast from "../../helpers/toast/errorToast";

const Attribute = ({
	attribute, fileId, setNeedRefresh, addNewAttribute, writeAccess,
}) => {
	const [clicked, setClicked] = useState(false);
	const [error, setError] = useState(false);

	const [attributeId, setAttributeId] = useState("");
	const [attributeName, setAttributeName] = useState("");
	const [value, setValue] = useState("");

	const handleClickOff = () => {
		setClicked(false);
		if (attributeName === "" || value === "") {
			setError(true);
		} else if (attribute.id === undefined) {
			addFileAttribute(fileId, attributeName, value, attribute.attribute_group_id).then((d) => {
				addNewAttribute();
				setAttributeId(d.id);
			}).catch((e) => {
				ErrorToast(e);
			});
		} else {
			updateFileAttribute(fileId, attributeId, attributeName, value, attribute.attribute_group_id).catch((e) => {
				ErrorToast(e);
			});
		}
	};

	const handleDelete = () => {
		deleteFileAttribute(fileId, attributeId).then(() => {
			setNeedRefresh(true);
		}).catch((e) => {
			ErrorToast(e);
		});
	};

	useEffect(() => {
		setAttributeId(attribute.id);
		setAttributeName(attribute.attribute_name);
		setValue(attribute.attribute_value);
	}, []);

	useEffect(() => {
		if (attributeName !== "" && value !== "") {
			setError(false);
		}
	}, [attributeName, value]);

	return (
		<tr>
			<td className="text-right font-light pr-2 pl-4 flex" aria-label="New attribute">
				<input
					className={`outline-none border-b-2 w-full text-right pl-2 mt-[2px] border-transparent
								${error ? "border-red-500 bg-red-100" : ""}
								${attributeName === "" || value === "" || clicked ? "border-gray-300" : ""}`}
					placeholder="New attribute"
					onFocus={() => { setClicked(true); }}
					onBlur={handleClickOff}
					value={attributeName}
					onChange={(e) => { setAttributeName(e.target.value); }}
					disabled={!writeAccess}
				/>:
			</td>
			<td className="font-medium" aria-label="New Value">
				<input
					className={`outline-none border-b-2 flex-auto px-2 border-transparent
								${error ? "border-red-500 bg-red-100" : ""}
								${attributeName === "" || value === "" || clicked ? "border-gray-300" : ""}`}
					placeholder="value"
					onFocus={() => { setClicked(true); }}
					onBlur={handleClickOff}
					value={value}
					onChange={(e) => { setValue(e.target.value); }}
					disabled={!writeAccess}
				/>
			</td>
			<td>
				{attributeId
					? (
						<button
							type="button"
							aria-label="delete"
							onClick={handleDelete}
							className={writeAccess ? "" : "hidden"}
							disabled={!writeAccess}
						>
							<TrashIcon className="h-6 w-6 text-gray-500 hover:text-red-500" />
						</button>
					)
					: (
						<button
							type="button"
							aria-label="save"
							className={writeAccess ? "" : "hidden"}
							disabled={!writeAccess}
						>
							<PlusIcon className="h-6 w-6 text-gray-300 hover:text-gray-500" />
						</button>
					)}
			</td>
		</tr>
	);
};

const Attributes = ({
	attributeGroup, fileId, setNeedRefresh, writeAccess,
}) => {
	const [attributeList, setAttributeList] = useState([]);
	const [collapsed, setCollapsed] = useState(true);

	const addNewAttribute = () => {
		setAttributeList([...attributeList, {
			attribute_group_id: attributeGroup.id,
			attribute_name: "",
			attribute_value: "",
		}]);
	};

	useEffect(() => {
		setAttributeList([...attributeGroup.attributes, {
			attribute_name: "",
			attribute_value: "",
			attribute_group_id: attributeGroup.id,
		}]);
	}, [attributeGroup]);

	return (
		<div className="w-full">
			<table className="w-full border-l-2 p-0 border-oxfordblue">
				<tbody>
					{attributeGroup.attribute_group_name !== "root"
						? (
							<tr>
								<th
									className="text-center font-semibold bg-oxfordblue text-offwhite pr-2"
									aria-label="Attribute"
									colSpan={2}
								>
									{attributeGroup.attribute_group_name}
								</th>
								<th
									className="text-right font-semibold bg-oxfordblue text-offwhite w-8"
									aria-label="Collapse"
								>
									<button type="button" onClick={() => { setCollapsed(!collapsed); }}>
										{collapsed
											? <ChevronDownIcon className="h-6 w-6 p-1 text-offwhite" />
											: (
												<ChevronDownIcon
													className="h-6 w-6 p-1 text-offwhite transform rotate-180"
												/>
											)}
									</button>
								</th>
							</tr>
						) : null}
					{!collapsed || attributeGroup.attribute_group_name === "root"
						? attributeList.map((attribute) => {
							return (
								<Attribute
									key={attribute.id ? attribute.id : "new"}
									attribute={attribute}
									fileId={fileId}
									setNeedRefresh={setNeedRefresh}
									addNewAttribute={addNewAttribute}
									writeAccess={writeAccess}
								/>
							);
						})
						: null }
					{(!collapsed || attributeGroup.attribute_group_name === "root")
						? attributeGroup.children.map((child) => {
							return (
								<tr>
									<td
										colSpan={3}
										className="border-l-2 p-0 border-oxfordblue"
										aria-label="Child Group"
									>
										<div className="flex">
											<div className="">
												<div className="h-10 w-4 bg-oxfordblue" />
											</div>
											<Attributes
												key={child.id}
												attributeGroup={child}
												fileId={fileId}
												setNeedRefresh={setNeedRefresh}
												writeAccess={writeAccess}
											/>
										</div>
									</td>
								</tr>
							);
						}) : null }
				</tbody>
			</table>
		</div>
	);
};

export default Attributes;
