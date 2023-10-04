#pragma once
#include <ydb/core/formats/arrow/replace_key.h>
#include <contrib/libs/apache/arrow/cpp/src/arrow/record_batch.h>

namespace NKikimr::NArrow {

class TSpecialKeys {
protected:
    std::shared_ptr<arrow::RecordBatch> Data;

    bool DeserializeFromString(const TString& data);

    std::optional<TReplaceKey> GetKeyByIndex(const ui32 position, const std::shared_ptr<arrow::Schema>& schema) const;

    TSpecialKeys() = default;
    TSpecialKeys(std::shared_ptr<arrow::RecordBatch> data)
        : Data(data) {
        Y_VERIFY(Data);
        Y_VERIFY(Data->num_rows());
    }

public:
    TString SerializeToStringDataOnlyNoCompression() const;

    TSpecialKeys(const TString& data, const std::shared_ptr<arrow::Schema>& schema) {
        Data = NArrow::DeserializeBatch(data, schema);
        Y_VERIFY(Data);
        Y_VERIFY_DEBUG(Data->ValidateFull().ok());
    }

    TSpecialKeys(const TString& data) {
        Y_VERIFY(DeserializeFromString(data));
    }

    TString SerializeToString() const;
};

class TFirstLastSpecialKeys: public TSpecialKeys {
private:
    using TBase = TSpecialKeys;
public:
    const std::shared_ptr<arrow::RecordBatch>& GetBatch() const {
        return Data;
    }

    std::shared_ptr<TFirstLastSpecialKeys> BuildAccordingToSchemaVerified(const std::shared_ptr<arrow::Schema>& schema) const;

    std::optional<TReplaceKey> GetFirst(const std::shared_ptr<arrow::Schema>& schema = nullptr) const {
        return GetKeyByIndex(0, schema);
    }
    std::optional<TReplaceKey> GetLast(const std::shared_ptr<arrow::Schema>& schema = nullptr) const {
        return GetKeyByIndex(Data->num_rows() - 1, schema);
    }

    explicit TFirstLastSpecialKeys(const TString& data);
    explicit TFirstLastSpecialKeys(const TString& data, const std::shared_ptr<arrow::Schema>& schema)
        : TBase(data, schema)
    {
        Y_VERIFY(Data->num_rows() == 1 || Data->num_rows() == 2);
    }
    explicit TFirstLastSpecialKeys(std::shared_ptr<arrow::RecordBatch> batch, const std::vector<TString>& columnNames = {});
};

class TMinMaxSpecialKeys: public TSpecialKeys {
private:
    using TBase = TSpecialKeys;
protected:
    TMinMaxSpecialKeys(std::shared_ptr<arrow::RecordBatch> data)
        : TBase(data) {
    }
public:
    std::shared_ptr<TMinMaxSpecialKeys> BuildAccordingToSchemaVerified(const std::shared_ptr<arrow::Schema>& schema) const;

    const std::shared_ptr<arrow::RecordBatch>& GetBatch() const {
        return Data;
    }

    std::optional<TReplaceKey> GetMin(const std::shared_ptr<arrow::Schema>& schema = nullptr) const {
        return GetKeyByIndex(0, schema);
    }
    std::optional<TReplaceKey> GetMax(const std::shared_ptr<arrow::Schema>& schema = nullptr) const {
        return GetKeyByIndex(Data->num_rows() - 1, schema);
    }

    explicit TMinMaxSpecialKeys(const TString& data);
    explicit TMinMaxSpecialKeys(const TString& data, const std::shared_ptr<arrow::Schema>& schema)
        : TBase(data, schema) {
        Y_VERIFY(Data->num_rows() == 1 || Data->num_rows() == 2);
    }

    explicit TMinMaxSpecialKeys(std::shared_ptr<arrow::RecordBatch> batch, const std::shared_ptr<arrow::Schema>& schema);
};

}
