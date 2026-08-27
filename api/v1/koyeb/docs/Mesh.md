# Mesh

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to [**MeshScope**](MeshScope.md) |  | [optional] [default to MESHSCOPE_UNSPECIFIED]
**Name** | Pointer to **string** | Custom mesh name — required when scope is MESH_SCOPE_CUSTOM, ignored otherwise. Combined with the workspace ID server-side to ensure the same custom name in different workspaces never collides. | [optional] 

## Methods

### NewMesh

`func NewMesh() *Mesh`

NewMesh instantiates a new Mesh object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeshWithDefaults

`func NewMeshWithDefaults() *Mesh`

NewMeshWithDefaults instantiates a new Mesh object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *Mesh) GetScope() MeshScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Mesh) GetScopeOk() (*MeshScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Mesh) SetScope(v MeshScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Mesh) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetName

`func (o *Mesh) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Mesh) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Mesh) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Mesh) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


