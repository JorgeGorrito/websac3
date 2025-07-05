import { AccessRequestForm } from "@/components/websac3/form/access-request/AccessRequestForm";
import { NoSessionLayout } from "@/layouts/NoSessionLayout";
import { ReactElement } from "react";


function AccessRequest() {
    return (
      <div className="absolute flex justify-center h-screen w-screen items-center">
        <div className="flex h-3/5 w-1/2 bg-white rounded-lg mt-12 shadow-2xl shadow-slate-800">
          <div className="flex flex-col w-full h-full pl-7 pr-7">
                <div className="flex justify-center items-center">
                    <h1 className="flex w-auto pl-4 pr-5 text-center font-bold text-gray-950 text-2xl mt-2 border-b-4 border-primary">Solicitud de Acceso</h1>
                </div>
                <div className="pt-5">
                    <AccessRequestForm />
                </div>
          </div>
        </div>
      </div>
    );
  }
  
  AccessRequest.getLayout = function getLayout(page: ReactElement) {
    return <NoSessionLayout>{page}</NoSessionLayout>;
  };
  
  export default AccessRequest;
  