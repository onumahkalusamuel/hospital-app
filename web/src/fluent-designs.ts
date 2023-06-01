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
  fluentCheckbox,
  fluentListbox
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
    fluentDialog(),
    fluentCheckbox(),
    fluentListbox()
  );
};

export default fluentDesigns;