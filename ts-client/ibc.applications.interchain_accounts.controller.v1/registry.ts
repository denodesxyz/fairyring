import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSendTx } from "./types/../../../../../pkg/mod/github.com/cosmos/ibc-go/v8@v8.2.1/proto/ibc/applications/interchain_accounts/controller/v1/tx";
import { QueryParamsResponse } from "./types/../../../../../pkg/mod/github.com/cosmos/ibc-go/v8@v8.2.1/proto/ibc/applications/interchain_accounts/controller/v1/query";
import { MsgRegisterInterchainAccount } from "./types/../../../../../pkg/mod/github.com/cosmos/ibc-go/v8@v8.2.1/proto/ibc/applications/interchain_accounts/controller/v1/tx";
import { MsgUpdateParams } from "./types/../../../../../pkg/mod/github.com/cosmos/ibc-go/v8@v8.2.1/proto/ibc/applications/interchain_accounts/controller/v1/tx";
import { MsgRegisterInterchainAccountResponse } from "./types/../../../../../pkg/mod/github.com/cosmos/ibc-go/v8@v8.2.1/proto/ibc/applications/interchain_accounts/controller/v1/tx";
import { MsgUpdateParamsResponse } from "./types/../../../../../pkg/mod/github.com/cosmos/ibc-go/v8@v8.2.1/proto/ibc/applications/interchain_accounts/controller/v1/tx";
import { QueryParamsRequest } from "./types/../../../../../pkg/mod/github.com/cosmos/ibc-go/v8@v8.2.1/proto/ibc/applications/interchain_accounts/controller/v1/query";
import { MsgSendTxResponse } from "./types/../../../../../pkg/mod/github.com/cosmos/ibc-go/v8@v8.2.1/proto/ibc/applications/interchain_accounts/controller/v1/tx";
import { Params } from "./types/../../../../../pkg/mod/github.com/cosmos/ibc-go/v8@v8.2.1/proto/ibc/applications/interchain_accounts/controller/v1/controller";
import { QueryInterchainAccountRequest } from "./types/../../../../../pkg/mod/github.com/cosmos/ibc-go/v8@v8.2.1/proto/ibc/applications/interchain_accounts/controller/v1/query";
import { QueryInterchainAccountResponse } from "./types/../../../../../pkg/mod/github.com/cosmos/ibc-go/v8@v8.2.1/proto/ibc/applications/interchain_accounts/controller/v1/query";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/ibc.applications.interchain_accounts.controller.v1.MsgSendTx", MsgSendTx],
    ["/ibc.applications.interchain_accounts.controller.v1.QueryParamsResponse", QueryParamsResponse],
    ["/ibc.applications.interchain_accounts.controller.v1.MsgRegisterInterchainAccount", MsgRegisterInterchainAccount],
    ["/ibc.applications.interchain_accounts.controller.v1.MsgUpdateParams", MsgUpdateParams],
    ["/ibc.applications.interchain_accounts.controller.v1.MsgRegisterInterchainAccountResponse", MsgRegisterInterchainAccountResponse],
    ["/ibc.applications.interchain_accounts.controller.v1.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/ibc.applications.interchain_accounts.controller.v1.QueryParamsRequest", QueryParamsRequest],
    ["/ibc.applications.interchain_accounts.controller.v1.MsgSendTxResponse", MsgSendTxResponse],
    ["/ibc.applications.interchain_accounts.controller.v1.Params", Params],
    ["/ibc.applications.interchain_accounts.controller.v1.QueryInterchainAccountRequest", QueryInterchainAccountRequest],
    ["/ibc.applications.interchain_accounts.controller.v1.QueryInterchainAccountResponse", QueryInterchainAccountResponse],
    
];

export { msgTypes }