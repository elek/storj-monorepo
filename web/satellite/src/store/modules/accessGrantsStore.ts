// Copyright (C) 2023 Storj Labs, Inc.
// See LICENSE for copying information.

import { defineStore } from 'pinia';
import { computed, reactive } from 'vue';

import {
    AccessGrant,
    AccessGrantCursor,
    AccessGrantsOrderBy,
    AccessGrantsPage,
    DurationPermission,
    EdgeCredentials,
} from '@/types/accessGrants';
import { SortDirection } from '@/types/common';
import { AccessGrantsApiGql } from '@/api/accessGrants';
import { useProjectsStore } from '@/store/modules/projectsStore';

class AccessGrantsState {
    public cursor: AccessGrantCursor = new AccessGrantCursor();
    public page: AccessGrantsPage = new AccessGrantsPage();
    public selectedAccessGrantsIds: string[] = [];
    public selectedBucketNames: string[] = [];
    public permissionNotBefore: Date | null = null;
    public permissionNotAfter: Date | null = null;
    public isDownload = true;
    public isUpload = true;
    public isList = true;
    public isDelete = true;
    public edgeCredentials: EdgeCredentials = new EdgeCredentials();
    public accessGrantsWebWorker: Worker | null = null;
    public isAccessGrantsWebWorkerReady = false;
}

export const useAccessGrantsStore = defineStore('accessGrants', () => {
    const api = new AccessGrantsApiGql();
    const { selectedProject } = useProjectsStore();

    const state = reactive<AccessGrantsState>(new AccessGrantsState());

    async function startWorker(): Promise<void> {
        const worker = new Worker(new URL('@/utils/accessGrant.worker.js', import.meta.url), { type: 'module' });
        worker.postMessage({ 'type': 'Setup' });

        const event: MessageEvent = await new Promise(resolve => worker.onmessage = resolve);
        if (event.data.error) {
            throw new Error(event.data.error);
        }

        if (event.data !== 'configured') {
            throw new Error('Failed to configure access grants web worker');
        }

        worker.onerror = (error: ErrorEvent) => {
            throw new Error(`Failed to configure access grants web worker. ${error.message}`);
        };

        state.accessGrantsWebWorker = worker;
        state.isAccessGrantsWebWorkerReady = true;
    }

    function stopWorker(): void {
        state.accessGrantsWebWorker?.terminate();
        state.accessGrantsWebWorker = null;
        state.isAccessGrantsWebWorkerReady = false;
    }

    async function fetchAccessGrants(pageNumber: number): Promise<AccessGrantsPage> {
        state.cursor.page = pageNumber;

        const accessGrantsPage: AccessGrantsPage = await api.get(selectedProject.id, state.cursor);

        state.page = accessGrantsPage;
        state.page.accessGrants = state.page.accessGrants.map(accessGrant => {
            if (state.selectedAccessGrantsIds.includes(accessGrant.id)) {
                accessGrant.isSelected = true;
            }

            return accessGrant;
        });

        return accessGrantsPage;
    }

    async function createAccessGrant(name: string): Promise<AccessGrant> {
        return await api.create(selectedProject.id, name);
    }

    async function deleteAccessGrants(): Promise<void> {
        await api.delete(state.selectedAccessGrantsIds);
    }

    async function deleteAccessGrantByNameAndProjectID(name: string): Promise<void> {
        await api.deleteByNameAndProjectID(name, selectedProject.id);
    }

    async function getEdgeCredentials(accessGrant: string, optionalURL: string, isPublic: boolean): Promise<EdgeCredentials> {
        const credentials: EdgeCredentials = await api.getGatewayCredentials(accessGrant, optionalURL, isPublic);

        state.edgeCredentials = credentials;

        return credentials;
    }

    function setAccessGrantsSearchQuery(query: string): void {
        state.cursor.search = query;
    }

    function setAccessGrantsSortingBy(order: AccessGrantsOrderBy): void {
        state.cursor.order = order;
    }

    function setAccessGrantsSortingDirection(direction: SortDirection): void {
        state.cursor.orderDirection = direction;
    }

    function setAccessGrantsDurationPermission(permission: DurationPermission): void {
        state.permissionNotBefore = permission.notBefore;
        state.permissionNotAfter = permission.notAfter;
    }

    function toggleAccessGrantsSelection(accessGrant: AccessGrant): void {
        if (!state.selectedAccessGrantsIds.includes(accessGrant.id)) {
            state.page.accessGrants.forEach((grant: AccessGrant) => {
                if (grant.id === accessGrant.id) {
                    grant.isSelected = true;
                }
            });
            state.selectedAccessGrantsIds.push(accessGrant.id);

            return;
        }

        state.page.accessGrants.forEach((grant: AccessGrant) => {
            if (grant.id === accessGrant.id) {
                grant.isSelected = false;
            }
        });
        state.selectedAccessGrantsIds = state.selectedAccessGrantsIds.filter(accessGrantId => {
            return accessGrant.id !== accessGrantId;
        });
    }

    function toggleBucketSelection(bucketName: string): void {
        if (!state.selectedBucketNames.includes(bucketName)) {
            state.selectedBucketNames.push(bucketName);

            return;
        }

        state.selectedBucketNames = state.selectedBucketNames.filter(name => {
            return bucketName !== name;
        });
    }

    function toggleIsDownloadPermission(): void {
        state.isDownload = !state.isDownload;
    }

    function toggleIsUploadPermission(): void {
        state.isUpload = !state.isUpload;
    }

    function toggleIsListPermission(): void {
        state.isList = !state.isList;
    }

    function toggleIsDeletePermission(): void {
        state.isDelete = !state.isDelete;
    }

    function clearAccessGrantsSelection(): void {
        state.selectedBucketNames = [];
        state.selectedAccessGrantsIds = [];
        state.page.accessGrants = state.page.accessGrants.map((accessGrant: AccessGrant) => {
            accessGrant.isSelected = false;

            return accessGrant;
        });
    }

    function clearAccessGrants(): void {
        state.cursor = new AccessGrantCursor();
        state.page = new AccessGrantsPage();
        state.selectedAccessGrantsIds = [];
        state.selectedBucketNames = [];
        state.permissionNotBefore = null;
        state.permissionNotAfter = null;
        state.edgeCredentials = new EdgeCredentials();
        state.isDownload = true;
        state.isUpload = true;
        state.isList = true;
        state.isDelete = true;
        state.accessGrantsWebWorker = null;
        state.isAccessGrantsWebWorkerReady = false;
    }

    const selectedAccessGrants = computed((): AccessGrant[] => {
        return state.page.accessGrants.filter((grant: AccessGrant) => grant.isSelected);
    });

    return {
        accessGrantsState: state,
        startWorker,
        stopWorker,
        fetchAccessGrants,
        createAccessGrant,
        deleteAccessGrants,
        deleteAccessGrantByNameAndProjectID,
        getEdgeCredentials,
        setAccessGrantsSearchQuery,
        setAccessGrantsSortingBy,
        setAccessGrantsSortingDirection,
        setAccessGrantsDurationPermission,
        toggleAccessGrantsSelection,
        toggleBucketSelection,
        toggleIsDownloadPermission,
        toggleIsUploadPermission,
        toggleIsListPermission,
        toggleIsDeletePermission,
        clearAccessGrantsSelection,
        clearAccessGrants,
        selectedAccessGrants,
    };
});
