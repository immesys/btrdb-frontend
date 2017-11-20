package main

import (
	"context"

	"github.com/immesys/smartgridstore/acl"
	"github.com/pborman/uuid"
)

const PRawValues = acl.ACLPermission("btrdb/RawValues")
const PAlignedWindows = acl.ACLPermission("btrdb/AlignedWindows")
const PWindows = acl.ACLPermission("btrdb/Windows")
const PStreamInfo = acl.ACLPermission("btrdb/StreamInfo")
const PSetStreamAnnotations = acl.ACLPermission("btrdb/SetStreamAnnotations")
const PChanges = acl.ACLPermission("btrdb/Changes")
const PCreate = acl.ACLPermission("btrdb/Create")
const PListCollections = acl.ACLPermission("btrdb/ListCollections")
const PLookupStreams = acl.ACLPermission("btrdb/LookupStreams")
const PNearest = acl.ACLPermission("btrdb/Nearest")
const PInsert = acl.ACLPermission("btrdb/Insert")
const PDelete = acl.ACLPermission("btrdb/Delete")
const PFlush = acl.ACLPermission("btrdb/Flush")
const PObliterate = acl.ACLPermission("btrdb/Obliterate")
const PFaultInject = acl.ACLPermission("btrdb/FaultInject")

type AccessControl struct {
}

func (ac *AccessControl) CheckPermissionsByUUID(ctx context.Context, uu uuid.UUID, op acl.ACLPermission) error {

}
