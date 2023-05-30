import {
  provideFluentDesignSystem,
  fluentCard,
  fluentButton,
  fluentTextField,
  fluentTextArea,
  fluentNumberField,
  fluentAnchor,
  fluentBreadcrumb,
  fluentBreadcrumbItem,
  fluentSkeleton,
  fluentRadio,
  fluentRadioGroup,
  fluentToolbar,
  fluentTabPanel,
  fluentDivider,
  fluentDialog,
} from '@fluentui/web-components';

const fluentDesigns = () => {
  provideFluentDesignSystem().register(
    fluentCard(),
    fluentButton(),
    fluentTextField(),
    fluentTextArea(),
    fluentNumberField(),
    fluentAnchor(),
    fluentBreadcrumb(),
    fluentBreadcrumbItem(),
    fluentSkeleton(),
    fluentRadio(),
    fluentRadioGroup(),
    fluentToolbar(),
    fluentTabPanel(),
    fluentDivider(),
    fluentDialog()
  );
};

export default fluentDesigns;