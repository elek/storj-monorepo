// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <CLIFlowContainer
        :on-back-click="onBackClick"
        :on-next-click="onNextClick"
        :is-loading="isLoading || areBucketNamesFetching"
        title="Access Permissions"
    >
        <template #icon>
            <Icon />
        </template>
        <template #content class="permissions">
            <div class="permissions__select">
                <p class="permissions__select__label">Select buckets to grant permission:</p>
                <VLoader v-if="areBucketNamesFetching" width="50px" height="50px" />
                <BucketsSelection v-else />
            </div>
            <div class="permissions__bucket-bullets">
                <div
                    v-for="(name, index) in selectedBucketNames"
                    :key="index"
                    class="permissions__bucket-bullets__container"
                >
                    <BucketNameBullet :name="name" />
                </div>
            </div>
            <div class="permissions__select">
                <p class="permissions__select__label">Choose permissions to allow:</p>
                <PermissionsSelect />
            </div>
            <div class="permissions__select">
                <p class="permissions__select__label">Duration of this access grant:</p>
                <DurationSelection />
            </div>
        </template>
    </CLIFlowContainer>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

import { RouteConfig } from "@/router";
import { BUCKET_ACTIONS } from "@/store/modules/buckets";
import { APP_STATE_MUTATIONS } from "@/store/mutationConstants";

import CLIFlowContainer from "@/components/onboardingTour/steps/common/CLIFlowContainer.vue";
import PermissionsSelect from "@/components/onboardingTour/steps/cliFlow/PermissionsSelect.vue";
import BucketNameBullet from "@/components/accessGrants/permissions/BucketNameBullet.vue";
import BucketsSelection from "@/components/accessGrants/permissions/BucketsSelection.vue";
import VLoader from "@/components/common/VLoader.vue";
import DurationSelection from "@/components/accessGrants/permissions/DurationSelection.vue";
import Icon from '@/../static/images/onboardingTour/accessGrant.svg';

// @vue/component
@Component({
    components: {
        CLIFlowContainer,
        PermissionsSelect,
        BucketNameBullet,
        BucketsSelection,
        VLoader,
        DurationSelection,
        Icon,
    }
})
export default class AGPermissions extends Vue {
    private worker: Worker;

    public areBucketNamesFetching = true;
    public isLoading = true;

    /**
     * Lifecycle hook after initial render.
     * Sets local key from props value.
     * Initializes web worker's onmessage functionality.
     */
    public async mounted(): Promise<void> {
        if (!this.cleanAPIKey) {
            this.isLoading = false;
            await this.onBackClick();

            return;
        }

        this.setWorker();

        try {
            await this.$store.dispatch(BUCKET_ACTIONS.FETCH_ALL_BUCKET_NAMES);

            this.areBucketNamesFetching = false;
        } catch (error) {
            await this.$notify.error(`Unable to fetch all bucket names. ${error.message}`);
        }

        this.isLoading = false;
    }

    /**
     * Sets local worker with worker instantiated in store.
     * Also sets worker's onmessage and onerror logic.
     */
    public setWorker(): void {
        this.worker = this.$store.state.accessGrantsModule.accessGrantsWebWorker;
        this.worker.onerror = (error: ErrorEvent) => {
            this.$notify.error(error.message);
        };
    }

    /**
     * Holds on back button click logic.
     * Navigates to previous screen.
     */
    public async onBackClick(): Promise<void> {
        if (this.isLoading) return;

        await this.$router.push({name: RouteConfig.AGName.name})
    }

    /**
     * Holds on next button click logic.
     */
    public async onNextClick(): Promise<void> {
        if (this.isLoading) return;

        this.isLoading = true;

        try {
            const restrictedKey = await this.generateRestrictedKey();
            await this.$store.commit(APP_STATE_MUTATIONS.SET_ONB_API_KEY, restrictedKey);
        } catch (error) {
            await this.$notify.error(error.message)
            this.isLoading = false;

            return;
        }

        this.isLoading = false;
        await this.$store.commit(APP_STATE_MUTATIONS.SET_ONB_API_KEY_STEP_BACK_ROUTE, this.$route.path)
        await this.$router.push({name: RouteConfig.APIKey.name});
    }

    /**
     * Generates and returns restricted key from clean API key.
     */
    private async generateRestrictedKey(): Promise<string> {
        let permissionsMsg = {
            'type': 'SetPermission',
            'isDownload': this.storedIsDownload,
            'isUpload': this.storedIsUpload,
            'isList': this.storedIsList,
            'isDelete': this.storedIsDelete,
            'buckets': this.selectedBucketNames,
            'apiKey': this.cleanAPIKey,
        }

        if (this.notBeforePermission) permissionsMsg = Object.assign(permissionsMsg, {'notBefore': this.notBeforePermission.toISOString()});
        if (this.notAfterPermission) permissionsMsg = Object.assign(permissionsMsg, {'notAfter': this.notAfterPermission.toISOString()});

        await this.worker.postMessage(permissionsMsg);

        const grantEvent: MessageEvent = await new Promise(resolve => this.worker.onmessage = resolve);
        if (grantEvent.data.error) {
            throw new Error(grantEvent.data.error)
        }

        return grantEvent.data.value;
    }

    /**
     * Returns selected bucket names.
     */
    public get selectedBucketNames(): string[] {
        return this.$store.state.accessGrantsModule.selectedBucketNames;
    }

    /**
     * Returns clean API key from store.
     */
    private get cleanAPIKey(): string {
        return this.$store.state.appStateModule.appState.onbCleanApiKey;
    }

    /**
     * Returns download permission from store.
     */
    private get storedIsDownload(): boolean {
        return this.$store.state.accessGrantsModule.isDownload;
    }

    /**
     * Returns upload permission from store.
     */
    private get storedIsUpload(): boolean {
        return this.$store.state.accessGrantsModule.isUpload;
    }

    /**
     * Returns list permission from store.
     */
    private get storedIsList(): boolean {
        return this.$store.state.accessGrantsModule.isList;
    }

    /**
     * Returns delete permission from store.
     */
    private get storedIsDelete(): boolean {
        return this.$store.state.accessGrantsModule.isDelete;
    }

    /**
     * Returns not before date permission from store.
     */
    private get notBeforePermission(): Date | null {
        return this.$store.state.accessGrantsModule.permissionNotBefore;
    }

    /**
     * Returns not after date permission from store.
     */
    private get notAfterPermission(): Date | null {
        return this.$store.state.accessGrantsModule.permissionNotAfter;
    }
}
</script>

<style scoped lang="scss">
    .permissions {

        &__select {

            &__label {
                font-size: 18px;
                line-height: 32px;
                letter-spacing: 0.75px;
                color: #4e4b66;
                margin: 20px 0 10px 0;
            }
        }

        &__bucket-bullets {
            display: flex;
            align-items: center;
            max-width: 100%;
            flex-wrap: wrap;

            &__container {
                display: flex;
                margin-top: 5px;
            }
        }
    }

    ::v-deep .buckets-selection,
    ::v-deep .duration-selection {
        margin-left: 0;
    }
</style>